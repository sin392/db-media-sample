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
		NewInput(name string) *FindShopByNameInput
		Execute(ctx context.Context, input *FindShopByNameInput) (*FindShopByNameOutput, error)
	}
	FindShopByNameInteractor struct {
		qRepo repository.ShopQueryRepository
	}
	FindShopByNameInput struct {
		name string
	}
	FindShopByNameOutput struct {
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

func (a *FindShopByNameInteractor) NewInput(name string) *FindShopByNameInput {
	return &FindShopByNameInput{
		name: name,
	}
}

func (a *FindShopByNameInteractor) newOutput(shop *model.Shop) *FindShopByNameOutput {
	return &FindShopByNameOutput{
		shop,
	}
}

func (a *FindShopByNameInteractor) Execute(ctx context.Context, input *FindShopByNameInput) (*FindShopByNameOutput, error) {
	ctx, span := trace.StartSpan(ctx, "FindShopByNameInteractor.Execute")
	defer span.End()

	shop, err := a.qRepo.FindByName(ctx, input.name)
	if err != nil {
		return nil, fmt.Errorf("failed to find shop by name: %w", err)
	}
	return a.newOutput(shop), nil
}
