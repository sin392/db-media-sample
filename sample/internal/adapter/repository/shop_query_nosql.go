package repository

import (
	"context"

	"github.com/sin392/db-media-sample/sample/internal/domain/model"
	"github.com/sin392/db-media-sample/sample/internal/domain/repository"
	"github.com/sin392/db-media-sample/sample/module/trace"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var _ repository.ShopQueryRepository = (*ShopQueryRepositoryNoSQLImpl)(nil)

const shopQueryCollectionName = "shops"

type ShopQueryRepositoryNoSQLImpl struct {
	db             NoSQL
	collectionName string
}

func NewShopQueryRepositoryNoSQL(db NoSQL) repository.ShopQueryRepository {
	return &ShopQueryRepositoryNoSQLImpl{
		db:             db,
		collectionName: shopQueryCollectionName,
	}
}

// FindByName 名前から店舗を取得する
func (r *ShopQueryRepositoryNoSQLImpl) FindByName(ctx context.Context, name string) (*model.Shop, error) {
	ctx, span := trace.StartSpan(ctx, "ShopRepositoryImpl.FindByName")
	defer span.End()

	var result model.Shop
	query := bson.M{
		// 大文字小文字を区別せずに部分一致検索
		"name": primitive.Regex{Pattern: name, Options: ""},
	}
	err := r.db.FindOne(ctx, r.collectionName, query, nil, &result)
	if err != nil {
		return nil, classifyError(err)
	}
	return &result, nil
}

// List 店舗一覧を取得する
func (r *ShopQueryRepositoryNoSQLImpl) List(ctx context.Context) (model.ShopList, error) {
	ctx, span := trace.StartSpan(ctx, "ShopRepositoryImpl.List")
	defer span.End()

	var results model.ShopList
	var query bson.M
	err := r.db.FindAll(ctx, r.collectionName, query, &results)
	if err != nil {
		return nil, classifyError(err)
	}
	return results, nil
}
