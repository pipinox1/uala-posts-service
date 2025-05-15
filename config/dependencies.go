package config

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"uala-posts-service/internal/domain/posts"
	"uala-posts-service/internal/domain/posts/content"
	"uala-posts-service/internal/infrastructure"
	"uala-posts-service/libs/events"
)

type Dependencies struct {
	PostRepository posts.Repository
	EventPublisher events.Publisher
	ContentFactory *content.ContentFactory
}

func BuildDependencies(config Config) (*Dependencies, error) {
	natsPublisher := events.NewNatsPublisher()
	url := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		config.Postgres.User,
		config.Postgres.Password,
		config.Postgres.Host,
		config.Postgres.Port,
		config.Postgres.Database,
	)
	if !config.Postgres.UseSSL {
		url += "?sslmode=disable"
	}
	db, err := sqlx.Connect("postgres", url)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	postRepository := infrastructure.NewPgPostRepository(db)
	contentFactory := content.NewContentFactory()
	return &Dependencies{
		PostRepository: postRepository,
		EventPublisher: natsPublisher,
		ContentFactory: contentFactory,
	}, nil
}
