package controller

import (
	"context"
	"fmt"

	"github.com/jinzhu/copier"
	"github.com/sin392/db-media-sample/sample/internal/usecase"
	"github.com/sin392/db-media-sample/sample/module/otel"
	pb "github.com/sin392/db-media-sample/sample/pb/shop/v1"
)

// バリデーションもここで行う.
func newFindShopByNameInput(req *pb.FindShopByNameRequest) (*usecase.FindShopByNameInput, error) {
	input := &usecase.FindShopByNameInput{
		Name: req.GetName(),
	}
	if err := input.Validate(); err != nil {
		return nil, fmt.Errorf("failed to validate input: %w", err)
	}
	return input, nil
}

func (c *ShopControllerPb) FindShopByName(ctx context.Context, req *pb.FindShopByNameRequest) (*pb.FindShopByNameResponse, error) {
	ctx, span := otel.StartSpan(ctx, "FindShopByNamePbController.FindShopByName")
	defer span.End()

	// リクエストのパースとバリデーション
	input, err := newFindShopByNameInput(req)
	if err != nil {
		return nil, fmt.Errorf("failed to create input: %w", err)
	}
	// usecaseの実行
	output, err := c.findShopByNameUc.Execute(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("failed to execute usecase: %w", err)
	}
	// レスポンス用の形式に変換
	var pbRes pb.FindShopByNameResponse
	if err := copier.Copy(&pbRes, output); err != nil {
		return nil, fmt.Errorf("failed to copy from output to pbRes: %w", err)
	}

	return &pbRes, nil
}
