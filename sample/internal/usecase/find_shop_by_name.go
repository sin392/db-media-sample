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
		Execute(ctx context.Context, input *FindShopByNameInput) (*FindShopByNameOutput, error)
	}
	FindShopByNameUsecaseImpl struct {
		qRepo repository.ShopQueryRepository
	}
	FindShopByNameInput struct {
		Name string
	}
	FindShopByNameOutput model.Shop
)

var _ FindShopByNameUsecase = (*FindShopByNameUsecaseImpl)(nil)

func (i *FindShopByNameInput) Validate() error {
	var err error
	if i.Name == "" {
		err = fmt.Errorf("name is required")
	}
	return err
}

func NewFindShopByNameUsecase(
	qRepo repository.ShopQueryRepository,
) FindShopByNameUsecase {
	return &FindShopByNameUsecaseImpl{
		qRepo: qRepo,
	}
}

func (a *FindShopByNameUsecaseImpl) newOutput(shop *model.Shop) *FindShopByNameOutput {
	output := FindShopByNameOutput(*shop)
	return &output
}

func (a *FindShopByNameUsecaseImpl) Execute(ctx context.Context, input *FindShopByNameInput) (*FindShopByNameOutput, error) {
	ctx, span := trace.StartSpan(ctx, "FindShopByNameUsecase.Execute")
	defer span.End()

	shop, err := a.qRepo.FindByName(ctx, input.Name)
	if err != nil {
		return nil, fmt.Errorf("failed to find shop by name: %w", err)
	}
	return a.newOutput(shop), nil
}
