package controller

import (
	"context"
	"fmt"

	"github.com/jinzhu/copier"
	"github.com/sin392/db-media-sample/sample/internal/adapter/presenter"
	"github.com/sin392/db-media-sample/sample/module/trace"
	pb "github.com/sin392/db-media-sample/sample/pb/shop/v1"
)

func (c *ShopPbController) FindShopByName(ctx context.Context, req *pb.FindShopByNameRequest) (*pb.FindShopByNameResponse, error) {
	ctx, span := trace.StartSpan(ctx, "FindShopByNamePbController.FindShopByName")
	defer span.End()

	// usecase用のInputに変換
	name := req.GetName()
	// usecaseの実行
	shop, err := c.findShopByNameUc.Execute(ctx, name)
	if err != nil {
		return nil, fmt.Errorf("failed to execute usecase: %w", err)
	}
	// プレゼンテーションのための変換
	output, err := presenter.OutputShopPassThrough(shop)
	if err != nil {
		return nil, fmt.Errorf("failed to pass through: %w", err)
	}
	// レスポンス用の形式に変換
	var pbRes pb.FindShopByNameResponse
	if err := copier.Copy(&pbRes, output); err != nil {
		return nil, fmt.Errorf("failed to copy from res to pbRes: %w", err)
	}

	return &pbRes, nil
}
