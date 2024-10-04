package controller

import (
	"context"
	"fmt"

	"github.com/jinzhu/copier"
	"github.com/sin392/db-media-sample/sample/module/trace"
	pb "github.com/sin392/db-media-sample/sample/pb/shop/v1"
)

func (c *ShopPbController) ListShop(ctx context.Context, req *pb.ListShopRequest) (*pb.ListShopResponse, error) {
	ctx, span := trace.StartSpan(ctx, "ShopPbController.ListShop")
	defer span.End()

	// usecaseの実行
	output, err := c.listShopUc.Execute(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to execute usecase: %w", err)
	}
	// レスポンス用の形式に変換
	var pbRes pb.ListShopResponse
	// ここのコピーもう少し改善できないか？
	if err := copier.Copy(&pbRes.Shops, output.ShopList); err != nil {
		return nil, fmt.Errorf("failed to copy from output to pbRes: %w", err)
	}

	return &pbRes, nil
}
