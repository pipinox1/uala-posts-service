package application

import (
	"context"
	"uala-posts-service/internal/domain/posts"
	"uala-posts-service/internal/domain/posts/content"
)

type UpdatePostContentCommand struct {
	PostID  string
	Content []ContentCommand
}

type UpdatePostContentResponse struct {
	Id string
}

type UpdatePostContent struct {
	contentFactory *content.ContentFactory
	postRepository posts.Repository
}

func NewUpdatePostContent(
	postRepository posts.Repository,
	contentFactory *content.ContentFactory,
) *UpdatePostContent {
	return &UpdatePostContent{
		postRepository: postRepository,
		contentFactory: contentFactory,
	}
}

func (s *UpdatePostContent) Exec(
	ctx context.Context,
	cmd *UpdatePostContentCommand,
) (*UpdatePostContentResponse, error) {
	post, err := s.postRepository.GetById(ctx, cmd.PostID)
	if err != nil {
		return nil, err
	}

	// TODO validate authorId with logged user
	contents := make([]content.Content, len(cmd.Content))

	// TODO we could use errgroup to parallelize the validation if needed
	for i, c := range cmd.Content {
		postContent, err := s.contentFactory.CreateContent(c.Type, content.ContentBody{
			Text: c.Text,
			Url:  c.Url,
		})
		if err != nil {
			return nil, err
		}
		contents[i] = *postContent
	}

	err = post.UpdatePostContent(contents)
	if err != nil {
		return nil, err
	}

	err = s.postRepository.Save(ctx, post)
	if err != nil {
		return nil, err
	}

	return &UpdatePostContentResponse{
		Id: post.ID,
	}, nil
}
