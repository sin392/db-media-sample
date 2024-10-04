package repository

import (
	"context"
	"fmt"

	"github.com/sin392/db-media-sample/sample/internal/domain/model"
	"github.com/sin392/db-media-sample/sample/internal/domain/repository"
	"github.com/sin392/db-media-sample/sample/module/trace"
)

var _ repository.ShopCommandRepository = (*ShopCommandRepositoryNoSQLImpl)(nil)

const shopCommandCollectionName = "shops"

type ShopCommandRepositoryNoSQLImpl struct {
	db             NoSQL
	collectionName string
}

func NewShopCommandRepositoryNoSQL(db NoSQL) repository.ShopCommandRepository {
	return &ShopCommandRepositoryNoSQLImpl{
		db:             db,
		collectionName: shopCommandCollectionName,
	}
}

func (r *ShopCommandRepositoryNoSQLImpl) Store(ctx context.Context, shop model.Shop) error {
	ctx, span := trace.StartSpan(ctx, "ShopCommandRepositoryNoSQLImpl.Store")
	defer span.End()

	if err := r.db.Store(ctx, r.collectionName, shop); err != nil {
		return fmt.Errorf("failed to store shop: %w", err)
	}
	return nil
}
