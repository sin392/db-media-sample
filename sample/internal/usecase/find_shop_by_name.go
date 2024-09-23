package usecase

import (
	"context"
	"time"

	"github.com/sin392/db-media-sample/internal/domain/model"
	"github.com/sin392/db-media-sample/internal/domain/repository"
)

type FindShopByNameUsecase interface {
	Execute(ctx context.Context, Name string) (*model.Shop, error)
}

type FindShopByNameOutput struct {
	Name   string  `json:"name"`
	Tel    string  `json:"tel"`
	Rating float32 `json:"rating"`
}

type findShopByNameInteractor struct {
	repo       repository.ShopRepository
	ctxTimeout time.Duration
}

func NewFindShopByNameIntercepter(
	repo repository.ShopRepository,
	ctxTimeout time.Duration,
) FindShopByNameUsecase {
	return &findShopByNameInteractor{
		repo:       repo,
		ctxTimeout: ctxTimeout,
	}
}

func (a *findShopByNameInteractor) Execute(ctx context.Context, Name string) (*model.Shop, error) {
	Shop, err := a.repo.FindByName(ctx, Name)
	if err != nil {
		return nil, err
	}
	return Shop, nil
}
