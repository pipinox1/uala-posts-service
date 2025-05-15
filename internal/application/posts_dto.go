package application

import "uala-posts-service/internal/domain/posts"

type PostDto struct {
	Id       string            `json:"id"`
	AuthorId string            `json:"author_id"`
	Content  []ContentResponse `json:"content"`
}

type ContentResponse struct {
	Type string  `json:"type"`
	Text *string `json:"text,omitempty"`
	Url  *string `json:"url,omitempty"`
}

func FromDomainToDto(post *posts.Post) *PostDto {
	contents := make([]ContentResponse, len(post.Contents))
	return &PostDto{
		Id:       post.ID,
		AuthorId: post.AuthorId,
		Content:  contents,
	}
}
