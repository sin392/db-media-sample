package repository

import (
	"context"

	"github.com/sin392/db-media-sample/sample/internal/domain/model"
)

type ShopRepository interface {
	FindByName(ctx context.Context, name string) (*model.Shop, error)
}
