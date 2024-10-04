package usecase

import (
	"context"
	"fmt"

	"github.com/sin392/db-media-sample/sample/internal/domain/model"
	"github.com/sin392/db-media-sample/sample/internal/domain/repository"
	"github.com/sin392/db-media-sample/sample/module/trace"
)

type (
	FindShopByNameUsecase interface {
		Execute(ctx context.Context, name string) (*ShopOutput, error)
	}
	FindShopByNameInteractor struct {
		repo repository.ShopRepository
	}
	// OutputData
	ShopOutput struct {
		*model.Shop
	}
)

func NewFindShopByNameIntercepter(
	repo repository.ShopRepository,
) FindShopByNameUsecase {
	return &FindShopByNameInteractor{
		repo: repo,
	}
}

func (a *FindShopByNameInteractor) Execute(ctx context.Context, name string) (*ShopOutput, error) {
	ctx, span := trace.StartSpan(ctx, "FindShopByNameInteractor.Execute")
	defer span.End()

	shop, err := a.repo.FindByName(ctx, name)
	if err != nil {
		return nil, fmt.Errorf("failed to find shop by name: %w", err)
	}
	res := &ShopOutput{
		shop,
	}
	return res, nil
}
