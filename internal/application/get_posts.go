package application

import (
	"context"
	"uala-posts-service/internal/domain/posts"
)

type GetPostsCommand struct {
	IDs []string
}

type GetPosts struct {
	postRepository posts.Repository
}

type GetPostsResponse struct {
	Posts []*PostDto `json:"posts"`
}

func NewGetPosts(postRepository posts.Repository) *GetPosts {
	return &GetPosts{
		postRepository: postRepository,
	}
}

func (s *GetPosts) Exec(ctx context.Context, cmd *GetPostsCommand) (*GetPostsResponse, error) {
	domainPosts, err := s.postRepository.MGetByIds(ctx, cmd.IDs)
	if err != nil {
		return nil, err
	}

	postsResponse := make([]*PostDto, len(domainPosts))
	for i, post := range domainPosts {
		postsResponse[i] = FromDomainToDto(post)
	}

	return &GetPostsResponse{
		Posts: postsResponse,
	}, nil

}
