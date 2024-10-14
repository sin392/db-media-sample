package controller

import (
	"context"
	"fmt"

	"github.com/jinzhu/copier"
	"github.com/sin392/db-media-sample/sample/internal/usecase"
	"github.com/sin392/db-media-sample/sample/module/otel"
	pb "github.com/sin392/db-media-sample/sample/pb/shop/v1"
)

// バリデーションもここで行う
func newStoreShopInput(req *pb.StoreShopRequest) (*usecase.StoreShopInput, error) {
	var input usecase.StoreShopInput
	if err := copier.Copy(&input, req); err != nil {
		return nil, fmt.Errorf("failed to convert from proto: %w", err)
	}
	if err := input.Validate(); err != nil {
		return nil, fmt.Errorf("failed to validate input: %w", err)
	}
	return &input, nil
}

func (c *ShopControllerPb) StoreShop(ctx context.Context, req *pb.StoreShopRequest) (*pb.StoreShopResponse, error) {
	ctx, span := otel.StartSpan(ctx, "StoreShopPbController.StoreShop")
	defer span.End()

	// リクエストのパースとバリデーション
	input, err := newStoreShopInput(req)
	if err != nil {
		return nil, fmt.Errorf("failed to create input: %w", err)
	}
	// usecaseの実行
	if err := c.storeShopUc.Execute(ctx, input); err != nil {
		return nil, fmt.Errorf("failed to execute usecase: %w", err)
	}

	return &pb.StoreShopResponse{}, nil
}
