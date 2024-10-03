package usecase

import (
	"context"

	"github.com/sin392/db-media-sample/sample/internal/domain/model"
	"github.com/sin392/db-media-sample/sample/internal/domain/repository"
	"github.com/sin392/db-media-sample/sample/module/trace"
)

type (
	FindShopByNameUsecase interface {
		Execute(ctx context.Context, name string) (*model.Shop, error)
	}
	FindShopByNameInteractor struct {
		repo repository.ShopRepository
	}
)

func NewFindShopByNameIntercepter(
	repo repository.ShopRepository,
) FindShopByNameUsecase {
	return &FindShopByNameInteractor{
		repo: repo,
	}
}

func (a *FindShopByNameInteractor) Execute(ctx context.Context, name string) (*model.Shop, error) {
	ctx, span := trace.StartSpan(ctx, "FindShopByNameInteractor.Execute")
	defer span.End()

	Shop, err := a.repo.FindByName(ctx, name)
	if err != nil {
		return nil, err
	}
	return Shop, nil
}
