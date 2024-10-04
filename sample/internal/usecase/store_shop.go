package usecase

import (
	"context"
	"fmt"

	"github.com/sin392/db-media-sample/sample/internal/domain/model"
	"github.com/sin392/db-media-sample/sample/internal/domain/repository"
	"github.com/sin392/db-media-sample/sample/module/trace"
)

type (
	StoreShopUsecase interface {
		Execute(ctx context.Context, input *StoreShopInput) error
	}
	StoreShopUsecaseImpl struct {
		qRepo repository.ShopCommandRepository
	}
	StoreShopInput struct {
		model.Shop
	}
)

var _ StoreShopUsecase = (*StoreShopUsecaseImpl)(nil)

// requiredのパラメータに関しては構造体作る段階でエラーが出るのでチェック不要
func (i *StoreShopInput) Validate() error {
	return nil
}

func NewStoreShopUsecase(
	qRepo repository.ShopCommandRepository,
) StoreShopUsecase {
	return &StoreShopUsecaseImpl{
		qRepo: qRepo,
	}
}

func (a *StoreShopUsecaseImpl) Execute(ctx context.Context, input *StoreShopInput) error {
	ctx, span := trace.StartSpan(ctx, "StoreShopUsecase.Execute")
	defer span.End()

	if err := a.qRepo.Store(ctx, input.Shop); err != nil {
		return fmt.Errorf("failed to store shop: %w", err)
	}
	return nil
}
