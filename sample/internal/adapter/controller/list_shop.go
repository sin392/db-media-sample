package controller

import (
	"context"
	"fmt"

	"github.com/sin392/db-media-sample/sample/internal/adapter/presenter"
	"github.com/sin392/db-media-sample/sample/internal/usecase"
	"github.com/sin392/db-media-sample/sample/module/trace"
)

type ListShopController struct {
	uc        usecase.ListShopUsecase
	presenter presenter.ListShopPresenter
}

func NewListShopController(uc usecase.ListShopUsecase, presenter presenter.ListShopPresenter) ListShopController {
	return ListShopController{
		uc:        uc,
		presenter: presenter,
	}
}

func (c *ListShopController) Execute(ctx context.Context) (*presenter.ListShopOutput, error) {
	ctx, span := trace.StartSpan(ctx, "ListShopController.Execute")
	defer span.End()

	output, err := c.uc.Execute(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to execute usecase: %w", err)
	}
	return c.presenter.Output(output), nil
}
