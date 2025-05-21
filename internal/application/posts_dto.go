package application

import (
	"time"
	"uala-posts-service/internal/domain/posts"
)

type PostDto struct {
	Id          string            `json:"id"`
	AuthorId    string            `json:"author_id"`
	Contents    []ContentResponse `json:"contents"`
	PublishedAt time.Time         `json:"published_at"`
	CreatedAt   time.Time         `json:"created_at"`
	UpdatedAt   time.Time         `json:"updated_at"`
}

type ContentResponse struct {
	Type string  `json:"type"`
	Text *string `json:"text,omitempty"`
	Url  *string `json:"url,omitempty"`
}

func FromDomainToDto(post *posts.Post) *PostDto {
	contents := make([]ContentResponse, len(post.Contents))
	for i, content := range post.Contents {
		contents[i] = ContentResponse{
			Type: content.Type.String(),
			Text: content.Text,
			Url:  content.Url,
		}
	}
	return &PostDto{
		Id:          post.ID,
		AuthorId:    post.AuthorId,
		Contents:    contents,
		PublishedAt: post.PublishedAt,
		CreatedAt:   post.CreatedAt,
		UpdatedAt:   post.UpdatedAt,
	}
}
