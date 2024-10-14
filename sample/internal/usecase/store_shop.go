package usecase

import (
	"context"
	"fmt"

	"github.com/sin392/db-media-sample/sample/internal/domain/model"
	"github.com/sin392/db-media-sample/sample/internal/domain/repository"
	appErrors "github.com/sin392/db-media-sample/sample/internal/errors"
	"github.com/sin392/db-media-sample/sample/module/snowflake"
	"github.com/sin392/db-media-sample/sample/module/trace"
)

type (
	StoreShopUsecase interface {
		Execute(ctx context.Context, input *StoreShopInput) error
	}
	StoreShopUsecaseImpl struct {
		qRepo repository.ShopCommandRepository
	}
	StoreShopInput model.Shop
)

var _ StoreShopUsecase = (*StoreShopUsecaseImpl)(nil)

func (i *StoreShopInput) Validate() error {
	var err error
	if i.Name == "" {
		err = fmt.Errorf("name is required")
	}
	if err != nil {
		return appErrors.NewApplicationError(appErrors.InvalidParameterError, err.Error())
	}
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

	// TODO: コンストラクタを使ってその中でIDを生成する
	shop := model.Shop(*input)
	shop.ID = snowflake.GetSnowflakeID(ctx)
	if err := a.qRepo.Store(ctx, shop); err != nil {
		return fmt.Errorf("failed to store shop: %w", err)
	}
	return nil
}
