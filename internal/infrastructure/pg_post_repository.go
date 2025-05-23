package infrastructure

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"time"
	"uala-posts-service/internal/domain/posts"
	"uala-posts-service/internal/domain/posts/content"
)

const (
	savePostQuery = `
        INSERT INTO posts (
            id,
            content,
            author_id,
            created_at,
            updated_at,
            deleted_at,
        	published_at
        ) VALUES (
            :id,
            :content,
            :author_id,
            :created_at,
            :updated_at,
            :deleted_at,
            :published_at
        )
    `

	getPostByIdQuery = `
        SELECT 
            id,
            content,
            author_id,
            created_at,
            updated_at,
            deleted_at,
            published_at
        FROM posts
        WHERE id = $1 AND deleted_at IS NULL
    `
	getPostByIdsQuery = `
        SELECT 
            id,
            content,
            author_id,
            created_at,
            updated_at,
            deleted_at,
        	published_at
        FROM posts
       WHERE id =any($1) AND deleted_at IS NULL
    `

	getPostsByAuthorIdQuery = `
        SELECT 
            id,
            content,
            author_id,
            created_at,
            updated_at,
            deleted_at,
            published_at
        FROM posts
        WHERE author_id = $1 AND deleted_at IS NULL
        ORDER BY created_at DESC
    `

	searchPostsBaseQuery = `
        SELECT 
            id,
            content,
            author_id,
            created_at,
            updated_at,
            deleted_at,
            published_at
        FROM posts
        WHERE deleted_at IS NULL
    `
)

var _ posts.Repository = (*PgPostRepository)(nil)

type PgPostRepository struct {
	db *sqlx.DB
}

func (i PgPostRepository) MGetByIds(ctx context.Context, ids []string) ([]*posts.Post, error) {
	var postsDB []postDB
	err := i.db.SelectContext(ctx, &postsDB, getPostByIdsQuery, pq.Array(ids))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return []*posts.Post{}, nil
		}
		return nil, posts.ErrPostInternalError
	}

	result := make([]*posts.Post, len(postsDB))
	for i, postDB := range postsDB {
		post, err := postDB.toDomain()
		if err != nil {
			return nil, fmt.Errorf("error converting post to domain: %w", err)
		}
		result[i] = post
	}

	return result, nil
}

func NewPgPostRepository(db *sqlx.DB) *PgPostRepository {
	return &PgPostRepository{db: db}
}

func (i PgPostRepository) Save(ctx context.Context, post *posts.Post) error {
	postDB := toPostDB(post)
	_, err := i.db.NamedExecContext(
		ctx,
		savePostQuery,
		postDB,
	)
	if err != nil {
		return fmt.Errorf("error saving post: %w", err)
	}

	return nil
}

func (i PgPostRepository) GetById(ctx context.Context, id string) (*posts.Post, error) {
	var postDB postDB
	err := i.db.GetContext(ctx, &postDB, getPostByIdQuery, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, posts.ErrPostNotFound
		}
		return nil, posts.ErrPostInternalError
	}

	return postDB.toDomain()
}

func (i PgPostRepository) GetByAuthorId(ctx context.Context, authorId string) ([]*posts.Post, error) {
	var postsDB []postDB
	err := i.db.SelectContext(ctx, &postsDB, getPostsByAuthorIdQuery, authorId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return []*posts.Post{}, nil
		}
		return nil, posts.ErrPostInternalError
	}

	result := make([]*posts.Post, 0, len(postsDB))
	for _, postDB := range postsDB {
		post, err := postDB.toDomain()
		if err != nil {
			return nil, fmt.Errorf("error converting post to domain: %w", err)
		}

		result = append(result, post)
	}

	return result, nil
}

type ContentJSON []ContentDB

func (c ContentJSON) Value() (driver.Value, error) {
	return json.Marshal(c)
}

func (c *ContentJSON) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &c)
}

type postDB struct {
	ID          string      `db:"id"`
	Content     ContentJSON `db:"content"`
	AuthorId    string      `db:"author_id"`
	CreatedAt   time.Time   `db:"created_at"`
	UpdatedAt   time.Time   `db:"updated_at"`
	DeletedAt   *time.Time  `db:"deleted_at"`
	PublishedAt time.Time   `db:"published_at"`
}

type ContentDB struct {
	Type string  `json:"type"`
	Text *string `json:"text"`
}

func toPostDB(post *posts.Post) postDB {
	contentJSON := make(ContentJSON, 0, len(post.Contents))
	for _, content := range post.Contents {
		contentJSON = append(contentJSON, ContentDB{
			Type: content.Type.String(),
			Text: content.Text,
		})
	}

	return postDB{
		ID:          post.ID,
		Content:     contentJSON,
		AuthorId:    post.AuthorId,
		CreatedAt:   post.CreatedAt,
		UpdatedAt:   post.UpdatedAt,
		DeletedAt:   post.DeletedAt,
		PublishedAt: post.PublishedAt,
	}
}

func (p postDB) toDomain() (*posts.Post, error) {
	contents := make([]content.Content, len(p.Content))
	for i, contentDB := range p.Content {
		contents[i] = content.Content{
			Type: content.ContentType(contentDB.Type),
			Text: contentDB.Text,
		}
	}

	return &posts.Post{
		ID:          p.ID,
		Contents:    contents,
		AuthorId:    p.AuthorId,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
		DeletedAt:   p.DeletedAt,
		PublishedAt: p.PublishedAt,
	}, nil
}
