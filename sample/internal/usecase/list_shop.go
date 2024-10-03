package usecase

import (
	"context"
	"fmt"

	"github.com/sin392/db-media-sample/sample/internal/domain/model"
	"github.com/sin392/db-media-sample/sample/internal/domain/repository"
	"github.com/sin392/db-media-sample/sample/module/trace"
)

type (
	ListShopUsecase interface {
		Execute(ctx context.Context) ([]model.Shop, error)
	}
	ListShopInteractor struct {
		repo repository.ShopRepository
	}
)

func NewListShopIntercepter(
	repo repository.ShopRepository,
) ListShopUsecase {
	return &ListShopInteractor{
		repo: repo,
	}
}

func (a *ListShopInteractor) Execute(ctx context.Context) ([]model.Shop, error) {
	ctx, span := trace.StartSpan(ctx, "ListShopInteractor.Execute")
	defer span.End()

	shops, err := a.repo.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to list shops: %w", err)
	}
	return shops, nil
}
