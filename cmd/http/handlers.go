package http

import (
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strings"
	"uala-posts-service/config"
	"uala-posts-service/internal/application"
	"uala-posts-service/internal/domain/posts"
)

var (
	MissingPostId = errors.New("missing post ids")
)

func createPost(deps *config.Dependencies) http.HandlerFunc {
	createPost := application.NewCreatePost(deps.PostRepository, deps.ContentFactory, deps.EventPublisher)
	return func(w http.ResponseWriter, r *http.Request) {
		var cmd application.CreatePostCommand
		err := json.NewDecoder(r.Body).Decode(&cmd)
		if err != nil {
			handleError(w, err)
			return
		}
		response, err := createPost.Exec(r.Context(), &cmd)
		if err != nil {
			handleError(w, err)
			return
		}

		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(response); err != nil {
			handleError(w, err)
			return
		}
	}
}

func getPost(deps *config.Dependencies) http.HandlerFunc {
	getPost := application.NewGetPostById(deps.PostRepository)
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "post_id")
		cmd := &application.GetPostByIdCommand{
			Id: id,
		}
		response, err := getPost.Exec(r.Context(), cmd)
		if err != nil {
			handleError(w, err)
			return
		}

		if err := json.NewEncoder(w).Encode(response); err != nil {
			handleError(w, err)
			return
		}
	}
}

func getPostByAuthor(deps *config.Dependencies) http.HandlerFunc {
	getPost := application.NewGetPostByAuthor(deps.PostRepository)
	return func(w http.ResponseWriter, r *http.Request) {
		authorID := chi.URLParam(r, "author_id")
		cmd := &application.GetPostsByAuthorCommand{
			AuthorID: authorID,
		}
		response, err := getPost.Exec(r.Context(), cmd)
		if err != nil {
			handleError(w, err)
			return
		}

		if err := json.NewEncoder(w).Encode(response); err != nil {
			handleError(w, err)
			return
		}
	}
}
func getPosts(deps *config.Dependencies) http.HandlerFunc {
	getPosts := application.NewGetPosts(deps.PostRepository)
	return func(w http.ResponseWriter, r *http.Request) {
		ids := r.URL.Query().Get("ids")
		if ids == "" {
			handleError(w, posts.ErrPostEmptyContent)
			return
		}
		cmd := &application.GetPostsCommand{
			IDs: strings.Split(ids, ","),
		}
		response, err := getPosts.Exec(r.Context(), cmd)
		if err != nil {
			handleError(w, err)
			return
		}

		if err := json.NewEncoder(w).Encode(response); err != nil {
			handleError(w, err)
			return
		}
	}
}

func handleError(w http.ResponseWriter, err error) {
	if err == nil {
		return
	}

	var errorResp ErrorResponse

	switch {
	case errors.Is(err, posts.ErrPostEmptyContent):
		errorResp = ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Empty Contents",
			Code:       err.Error(),
		}
	case errors.Is(err, MissingPostId):
		errorResp = ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Missing post ids",
			Code:       err.Error(),
		}
	default:
		errorResp = ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "Internal server error",
			Code:       "INTERNAL_ERROR",
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(errorResp.StatusCode)

	jsonResp, jsonErr := json.Marshal(errorResp)
	if jsonErr != nil {
		http.Error(w, "Error processing response", http.StatusInternalServerError)
		return
	}

	w.Write(jsonResp)
	return
}

type ErrorResponse struct {
	StatusCode int    `json:"status,omitempty"`
	Message    string `json:"message,omitempty"`
	Code       string `json:"code,omitempty"`
}
