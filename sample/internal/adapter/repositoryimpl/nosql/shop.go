package nosql

import (
	"context"

	"github.com/sin392/db-media-sample/internal/domain/model"
	"github.com/sin392/db-media-sample/internal/domain/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var _ repository.ShopRepository = (*ShopRepositoryImpl)(nil)

type ShopRepositoryImpl struct {
	db             NoSQL
	collectionName string
}

func NewShopRepositoryImpl(db NoSQL) repository.ShopRepository {
	return &ShopRepositoryImpl{
		db:             db,
		collectionName: "shops",
	}
}

func (r *ShopRepositoryImpl) FindByName(ctx context.Context, name string) (*model.Shop, error) {
	var result model.Shop
	query := bson.M{
		// 大文字小文字を区別せずに部分一致検索
		"name": primitive.Regex{Pattern: name, Options: ""},
	}
	err := r.db.FindOne(ctx, r.collectionName, query, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
