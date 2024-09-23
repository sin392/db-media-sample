package usecase

import (
	"context"
	"time"

	"github.com/sin392/db-media-sample/internal/domain/model"
	"github.com/sin392/db-media-sample/internal/domain/repository"
	"github.com/sin392/db-media-sample/module/trace"
)

type FindShopByNameUsecase interface {
	Execute(ctx context.Context, name string) (*model.Shop, error)
}

type FindShopByNameOutput struct {
	model.Shop
}

type FindShopByNameInteractor struct {
	repo       repository.ShopRepository
	ctxTimeout time.Duration
}

func NewFindShopByNameIntercepter(
	repo repository.ShopRepository,
	ctxTimeout time.Duration,
) FindShopByNameUsecase {
	return &FindShopByNameInteractor{
		repo:       repo,
		ctxTimeout: ctxTimeout,
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
