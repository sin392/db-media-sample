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
		Execute(ctx context.Context) (*ShopListOutput, error)
	}
	ListShopInteractor struct {
		qRepo repository.ShopQueryRepository
	}
	ShopListOutput struct {
		model.ShopList
	}
)

func NewListShopIntercepter(
	qRepo repository.ShopQueryRepository,
) ListShopUsecase {
	return &ListShopInteractor{
		qRepo: qRepo,
	}
}

func (a *ListShopInteractor) newOutput(shops model.ShopList) *ShopListOutput {
	return &ShopListOutput{
		shops,
	}
}

func (a *ListShopInteractor) Execute(ctx context.Context) (*ShopListOutput, error) {
	ctx, span := trace.StartSpan(ctx, "ListShopInteractor.Execute")
	defer span.End()

	shops, err := a.qRepo.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to list shops: %w", err)
	}
	return a.newOutput(shops), nil
}
