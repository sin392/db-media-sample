package usecase

import (
	"context"
	"fmt"

	"github.com/sin392/db-media-sample/sample/internal/domain/model"
	"github.com/sin392/db-media-sample/sample/internal/domain/repository"
	"github.com/sin392/db-media-sample/sample/module/otel"
)

type (
	ListShopUsecase interface {
		Execute(ctx context.Context) (*ShopListOutput, error)
	}
	ListShopUsecaseImpl struct {
		qRepo repository.ShopQueryRepository
	}
	ShopListOutput model.ShopList
)

var _ ListShopUsecase = (*ListShopUsecaseImpl)(nil)

func NewListShopUsecase(
	qRepo repository.ShopQueryRepository,
) ListShopUsecase {
	return &ListShopUsecaseImpl{
		qRepo: qRepo,
	}
}

func (a *ListShopUsecaseImpl) newOutput(shops model.ShopList) *ShopListOutput {
	output := ShopListOutput(shops)

	return &output
}

func (a *ListShopUsecaseImpl) Execute(ctx context.Context) (*ShopListOutput, error) {
	ctx, span := otel.StartSpan(ctx, "ListShopUsecaseImpl.Execute")
	defer span.End()

	shops, err := a.qRepo.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to list shops: %w", err)
	}

	return a.newOutput(shops), nil
}
