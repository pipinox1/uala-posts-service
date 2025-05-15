package http

import (
	"encoding/json"
	"errors"
	"net/http"
	"uala-posts-service/config"
	"uala-posts-service/internal/application"
	"uala-posts-service/internal/domain/posts"
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
