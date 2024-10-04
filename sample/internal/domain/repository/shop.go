package repository

import (
	"context"

	"github.com/sin392/db-media-sample/sample/internal/domain/model"
)

type ShopRepository interface {
	// FindByName 名前から店舗を取得する
	FindByName(ctx context.Context, name string) (*model.Shop, error)
	// List 店舗一覧を取得する
	List(ctx context.Context) (model.ShopList, error)
}
