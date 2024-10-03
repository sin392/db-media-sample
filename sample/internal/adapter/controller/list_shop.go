package controller

import (
	"context"
	"fmt"

	"github.com/sin392/db-media-sample/sample/internal/domain/model"
	"github.com/sin392/db-media-sample/sample/internal/usecase"
	"github.com/sin392/db-media-sample/sample/module/trace"
)

type ListShopController struct {
	uc usecase.ListShopUsecase
}

func NewListShopController(uc usecase.ListShopUsecase) ListShopController {
	return ListShopController{
		uc: uc,
	}
}

func (c *ListShopController) Execute(ctx context.Context) ([]model.Shop, error) {
	ctx, span := trace.StartSpan(ctx, "ListShopController.Execute")
	defer span.End()

	shops, err := c.uc.Execute(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to execute usecase: %w", err)
	}
	return shops, nil
}
