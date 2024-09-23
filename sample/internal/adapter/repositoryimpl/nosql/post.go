package nosql

import (
	"context"

	"github.com/sin392/db-media-sample/internal/domain/model"
	"github.com/sin392/db-media-sample/internal/domain/repository"
	"go.mongodb.org/mongo-driver/bson"
)

var _ repository.PostRepository = (*PostRepositoryImpl)(nil)

const tableName = "posts"

type PostRepositoryImpl struct {
	db NoSQL
}

func NewPostRepositoryImpl(db NoSQL) repository.PostRepository {
	return &PostRepositoryImpl{
		db: db,
	}
}

func (r *PostRepositoryImpl) FindByTitle(ctx context.Context, title string) (*model.Post, error) {
	var result model.Post
	query := bson.M{
		"title": title,
	}
	err := r.db.FindOne(ctx, tableName, query, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
