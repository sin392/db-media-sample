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
		qRepo repository.ShopQueryRepository
	}
	// OutputData
	ShopOutput struct {
		*model.Shop
	}
)

func NewFindShopByNameIntercepter(
	qRepo repository.ShopQueryRepository,
) FindShopByNameUsecase {
	return &FindShopByNameInteractor{
		qRepo: qRepo,
	}
}

func (a *FindShopByNameInteractor) Execute(ctx context.Context, name string) (*ShopOutput, error) {
	ctx, span := trace.StartSpan(ctx, "FindShopByNameInteractor.Execute")
	defer span.End()

	shop, err := a.qRepo.FindByName(ctx, name)
	if err != nil {
		return nil, fmt.Errorf("failed to find shop by name: %w", err)
	}
	res := &ShopOutput{
		shop,
	}
	return res, nil
}
