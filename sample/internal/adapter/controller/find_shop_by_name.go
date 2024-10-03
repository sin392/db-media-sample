package controller

import (
	"context"
	"fmt"

	"github.com/sin392/db-media-sample/sample/internal/adapter/presenter"
	"github.com/sin392/db-media-sample/sample/internal/usecase"
	"github.com/sin392/db-media-sample/sample/module/trace"
)

type FindShopByNameController struct {
	uc        usecase.FindShopByNameUsecase
	presenter presenter.FindShopByNamePresenter
}

func NewFindShopByNameController(uc usecase.FindShopByNameUsecase, presenter presenter.FindShopByNamePresenter) FindShopByNameController {
	return FindShopByNameController{
		uc:        uc,
		presenter: presenter,
	}
}

func (c *FindShopByNameController) Execute(ctx context.Context, name string) (*presenter.FindShopByNameOutput, error) {
	ctx, span := trace.StartSpan(ctx, "FindShopByNameController.Execute")
	defer span.End()

	output, err := c.uc.Execute(ctx, name)
	if err != nil {
		return nil, fmt.Errorf("failed to execute usecase: %w", err)
	}
	return c.presenter.Output(output), nil
}
