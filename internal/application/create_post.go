package application

import (
	"context"
	"uala-posts-service/internal/domain/posts"
	"uala-posts-service/internal/domain/posts/content"
	"uala-posts-service/libs/events"
)

type CreatePostCommand struct {
	AuthorId string           `json:"author_id"`
	Contents []ContentCommand `json:"contents"`
}

type ContentCommand struct {
	Type string  `json:"type"`
	Text *string `json:"text"`
	Url  *string `json:"url"`
}

type CreatePostResponse struct {
	*PostDto
}

type CreatePost struct {
	contentFactory *content.ContentFactory
	postRepository posts.Repository
	eventPublisher events.Publisher
}

func NewCreatePost(
	postRepository posts.Repository,
	contentFactory *content.ContentFactory,
	eventPublisher events.Publisher,
) *CreatePost {
	return &CreatePost{
		postRepository: postRepository,
		contentFactory: contentFactory,
		eventPublisher: eventPublisher,
	}
}

func (s *CreatePost) Exec(
	ctx context.Context,
	cmd *CreatePostCommand,
) (*CreatePostResponse, error) {
	// TODO Validate user
	/*
		user, err := auth.GetUserFromContext(ctx)
		if err != nil {
			return nil, err
		}

		if user != cmd.AuthorId {
			return nil, domain.InvalidUsers
		}
	*/

	// TODO we could use errgroup to parallelize the validation if needed
	contents := make([]content.Content, len(cmd.Contents))
	for i, c := range cmd.Contents {
		postContent, err := s.contentFactory.CreateContent(c.Type, content.ContentBody{
			Text: c.Text,
			Url:  c.Url,
		})
		if err != nil {
			return nil, err
		}
		contents[i] = *postContent
	}

	post, err := posts.CreatePost(cmd.AuthorId, contents)
	if err != nil {
		return nil, err
	}

	err = s.postRepository.Save(ctx, post)
	if err != nil {
		return nil, err
	}

	go func() {
		err := s.eventPublisher.Publish(context.WithoutCancel(ctx), posts.NewPostCreatedEvent(post))
		if err != nil {
			//TODO log error and metric
		}
	}()

	return &CreatePostResponse{
		FromDomainToDto(post),
	}, nil
}
