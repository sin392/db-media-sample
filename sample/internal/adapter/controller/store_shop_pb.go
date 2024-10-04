package controller

import (
	"context"
	"fmt"

	"github.com/sin392/db-media-sample/sample/internal/domain/model"
	"github.com/sin392/db-media-sample/sample/internal/usecase"
	"github.com/sin392/db-media-sample/sample/module/trace"
	pb "github.com/sin392/db-media-sample/sample/pb/shop/v1"
)

// バリデーションもここで行う
func newStoreShopInput(req *pb.StoreShopRequest) (*usecase.StoreShopInput, error) {
	input := &usecase.StoreShopInput{
		// マッピングくそだるくないか？
		// ドメインモデル埋め込んじゃうとコンストラクタ経由以外でドメインモデルできちゃうのも気になる
		Shop: model.Shop{
			Name: req.GetName(),
			Location: model.Location{
				Prefecture: req.GetLocation().GetPrefecture(),
				City:       req.GetLocation().GetCity(),
				Address:    req.GetLocation().GetAddress(),
			},
			// TODO...
		},
	}
	if err := input.Validate(); err != nil {
		return nil, fmt.Errorf("failed to validate input: %w", err)
	}
	return input, nil
}

func (c *ShopControllerPb) StoreShop(ctx context.Context, req *pb.StoreShopRequest) (*pb.StoreShopResponse, error) {
	ctx, span := trace.StartSpan(ctx, "StoreShopPbController.StoreShop")
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
