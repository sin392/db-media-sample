package controller

import (
	"context"
	"fmt"

	"github.com/sin392/db-media-sample/sample/internal/domain/model"
	"github.com/sin392/db-media-sample/sample/internal/usecase"
	"github.com/sin392/db-media-sample/sample/module/trace"
)

type FindShopByNameController struct {
	uc usecase.FindShopByNameUsecase
}

func NewFindShopByNameController(uc usecase.FindShopByNameUsecase) FindShopByNameController {
	return FindShopByNameController{
		uc: uc,
	}
}

func (c *FindShopByNameController) Execute(ctx context.Context, name string) (*model.Shop, error) {
	ctx, span := trace.StartSpan(ctx, "FindShopByNameController.Execute")
	defer span.End()

	shop, err := c.uc.Execute(ctx, name)
	if err != nil {
		return nil, fmt.Errorf("failed to execute usecase: %w", err)
	}
	return shop, nil
}
