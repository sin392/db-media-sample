package nosql

import (
	"context"
	"fmt"

	"github.com/sin392/db-media-sample/sample/internal/domain/model"
	"github.com/sin392/db-media-sample/sample/internal/domain/repository"
	"github.com/sin392/db-media-sample/sample/module/trace"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var _ repository.ShopRepository = (*ShopRepositoryImpl)(nil)

const shopCollectionName = "shops"

type ShopRepositoryImpl struct {
	db             NoSQL
	collectionName string
}

func NewShopRepositoryImpl(db NoSQL) repository.ShopRepository {
	return &ShopRepositoryImpl{
		db:             db,
		collectionName: shopCollectionName,
	}
}

// FindByName 名前から店舗を取得する
func (r *ShopRepositoryImpl) FindByName(ctx context.Context, name string) (*model.Shop, error) {
	ctx, span := trace.StartSpan(ctx, "ShopRepositoryImpl.FindByName")
	defer span.End()

	var result model.Shop
	query := bson.M{
		// 大文字小文字を区別せずに部分一致検索
		"name": primitive.Regex{Pattern: name, Options: ""},
	}
	err := r.db.FindOne(ctx, r.collectionName, query, nil, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to find shop by name: %w", err)
	}
	return &result, nil
}

// List 店舗一覧を取得する
func (r *ShopRepositoryImpl) List(ctx context.Context) ([]model.Shop, error) {
	ctx, span := trace.StartSpan(ctx, "ShopRepositoryImpl.List")
	defer span.End()

	var results []model.Shop
	query := bson.M{}
	err := r.db.FindAll(ctx, r.collectionName, query, &results)
	if err != nil {
		return nil, fmt.Errorf("failed to find all shops: %w", err)
	}
	return results, nil
}
