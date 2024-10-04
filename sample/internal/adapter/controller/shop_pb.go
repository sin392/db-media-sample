package controller

import (
	"github.com/sin392/db-media-sample/sample/internal/usecase"
	pb "github.com/sin392/db-media-sample/sample/pb/shop/v1"
)

type ShopPbController struct {
	pb.UnimplementedShopServiceServer
	findShopByNameUc usecase.FindShopByNameUsecase
	listShopUc       usecase.ListShopUsecase
}

var _ pb.ShopServiceServer = (*ShopPbController)(nil)

// pb.ShopServiceServerを実装したShopPbControllerを生成する
func NewShopPbController(
	findShopByNameUc usecase.FindShopByNameUsecase,
	listShopUc usecase.ListShopUsecase,
) ShopPbController {
	return ShopPbController{
		findShopByNameUc: findShopByNameUc,
		listShopUc:       listShopUc,
	}
}
