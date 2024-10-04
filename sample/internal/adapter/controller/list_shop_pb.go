package controller

import (
	"context"
	"fmt"

	"github.com/jinzhu/copier"
	"github.com/sin392/db-media-sample/sample/internal/adapter/presenter"
	"github.com/sin392/db-media-sample/sample/module/trace"
	pb "github.com/sin392/db-media-sample/sample/pb/shop/v1"
)

func (c *ShopPbController) ListShop(ctx context.Context, req *pb.ListShopRequest) (*pb.ListShopResponse, error) {
	ctx, span := trace.StartSpan(ctx, "ShopPbController.ListShop")
	defer span.End()

	// usecaseの実行
	shops, err := c.listShopUc.Execute(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to execute usecase: %w", err)
	}
	// プレゼンテーションのための変換
	output, err := presenter.OutputShopListPassThrough(shops)
	if err != nil {
		return nil, fmt.Errorf("failed to pass through: %w", err)
	}
	// レスポンス用の形式に変換
	var pbRes pb.ListShopResponse
	// ここのコピーもう少し改善できないか？
	if err := copier.Copy(&pbRes.Shops, output.ShopListOutput.ShopList); err != nil {
		return nil, fmt.Errorf("failed to copy from res to pbRes: %w", err)
	}

	return &pbRes, nil
}
