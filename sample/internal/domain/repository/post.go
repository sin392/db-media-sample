package repository

import (
	"context"

	"github.com/sin392/db-media-sample/internal/domain/model"
)

type PostRepository interface {
	FindByTitle(ctx context.Context, title string) (*model.Post, error)
	// FindAll(ctx context.Context) ([]*model.Post, error)
}
