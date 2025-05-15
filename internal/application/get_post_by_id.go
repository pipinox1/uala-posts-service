package application

import (
	"context"
	"uala-posts-service/internal/domain/posts"
)

type GetPostByIdCommand struct {
	Id string
}

type GetPostByIdResponse struct {
	*PostDto
}

type GetPostById struct {
	postRepository posts.Repository
}

func NewGetPostById(postRepository posts.Repository) *GetPostById {
	return &GetPostById{
		postRepository: postRepository,
	}
}

func (s *GetPostById) Exec(ctx context.Context, cmd *GetPostByIdCommand) (*GetPostByIdResponse, error) {
	post, err := s.postRepository.GetById(ctx, cmd.Id)
	if err != nil {
		return nil, err
	}

	return &GetPostByIdResponse{
		FromDomainToDto(post),
	}, nil
}
