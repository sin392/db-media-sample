package repository

import (
	"context"

	"github.com/sin392/db-media-sample/sample/internal/domain/model"
)

type ShopQueryRepository interface {
	// FindByName 名前から店舗を取得する
	FindByName(ctx context.Context, name string) (*model.Shop, error)
	// List 店舗一覧を取得する
	List(ctx context.Context) (model.ShopList, error)
}

type ShopCommandRepository interface {
	// Store 店舗を保存する
	Store(ctx context.Context, shop model.Shop) error
}
