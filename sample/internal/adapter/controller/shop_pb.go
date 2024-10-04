package controller

import (
	"github.com/sin392/db-media-sample/sample/internal/usecase"
	pb "github.com/sin392/db-media-sample/sample/pb/shop/v1"
)

type ShopControllerPb struct {
	pb.UnimplementedShopServiceServer
	findShopByNameUc usecase.FindShopByNameUsecase
	listShopUc       usecase.ListShopUsecase
}

var _ pb.ShopServiceServer = (*ShopControllerPb)(nil)

// pb.ShopServiceServerを実装したShopControllerPbを生成する
func NewShopControllerPb(
	findShopByNameUc usecase.FindShopByNameUsecase,
	listShopUc usecase.ListShopUsecase,
) ShopControllerPb {
	return ShopControllerPb{
		findShopByNameUc: findShopByNameUc,
		listShopUc:       listShopUc,
	}
}
