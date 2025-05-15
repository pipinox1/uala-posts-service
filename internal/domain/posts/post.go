package posts

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"time"
	"uala-posts-service/internal/domain/posts/content"
)

var ErrPostEmptyContent = errors.New("post.empty_content")
var ErrEmptyAuthorId = errors.New("post.empty_author_id")

type Repository interface {
	Save(ctx context.Context, post *Post) error
	GetById(ctx context.Context, id string) (*Post, error)
	GetByAuthorId(ctx context.Context, authorId string) ([]*Post, error)
}

type Post struct {
	ID        string
	Contents  []content.Content
	AuthorId  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type PostSearchFilters struct {
	Offset *int
	Limit  *int
}

func CreatePost(
	authorId string,
	content []content.Content,
) (*Post, error) {
	if len(content) == 0 {
		return nil, ErrPostEmptyContent
	}
	if authorId == "" {
		return nil, ErrEmptyAuthorId
	}

	now := time.Now()

	return &Post{
		ID:        uuid.New().String(),
		Contents:  content,
		AuthorId:  authorId,
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}

func (p *Post) UpdatePostContent(
	content []content.Content,
) error {
	p.Contents = content
	p.UpdatedAt = time.Now()
	return nil
}
