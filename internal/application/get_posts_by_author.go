package application

import (
	"context"
	"uala-posts-service/internal/domain/posts"
)

type GetPostsByAuthorCommand struct {
	AuthorID string
}

type GetPostsByAuthorResponse struct {
	Posts []*PostDto `json:"posts"`
}

type GetPostByAuthor struct {
	postRepository posts.Repository
}

func NewGetPostByAuthor(postRepository posts.Repository) *GetPostByAuthor {
	return &GetPostByAuthor{
		postRepository: postRepository,
	}
}

func (s *GetPostByAuthor) Exec(ctx context.Context, cmd *GetPostsByAuthorCommand) (*GetPostsByAuthorResponse, error) {
	domainPosts, err := s.postRepository.GetByAuthorId(ctx, cmd.AuthorID)
	if err != nil {
		return nil, err
	}

	postsResponse := make([]*PostDto, len(domainPosts))
	for i, post := range domainPosts {
		postsResponse[i] = FromDomainToDto(post)
	}

	return &GetPostsByAuthorResponse{
		Posts: postsResponse,
	}, nil
}
