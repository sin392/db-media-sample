package router

import (
	"context"
	"fmt"

	"github.com/jinzhu/copier"
	"github.com/sin392/db-media-sample/sample/internal/adapter/controller"
	pb "github.com/sin392/db-media-sample/sample/pb/shop/v1"
	"google.golang.org/grpc"
)

type ShopRouter struct {
	pb.ShopServiceServer
	findShopByNameCtrl controller.FindShopByNameController
	listShopCtrl       controller.ListShopController
}

func NewShopRouter(
	findShopByNameCtrl controller.FindShopByNameController,
	listShopCtrl controller.ListShopController,
) ShopRouter {
	return ShopRouter{
		findShopByNameCtrl: findShopByNameCtrl,
		listShopCtrl:       listShopCtrl,
	}
}

func (r *ShopRouter) FindShopByName(ctx context.Context, req *pb.FindShopByNameRequest) (*pb.FindShopByNameResponse, error) {
	name := req.GetName()
	shop, err := r.findShopByNameCtrl.Execute(ctx, name)
	if err != nil {
		return nil, fmt.Errorf("failed to execute controller: %w", err)
	}
	menus := make([]*pb.Menu, len(shop.Menus))
	for i, m := range shop.Menus {
		menus[i] = &pb.Menu{
			Name:  m.Name,
			Price: int32(m.Price),
			Desc:  m.Desc,
		}
	}
	var res pb.FindShopByNameResponse
	copier.Copy(&res, shop)
	return &res, nil
}

func (r *ShopRouter) ListShop(ctx context.Context, req *pb.ListShopRequest) (*pb.ListShopResponse, error) {
	shops, err := r.listShopCtrl.Execute(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to execute controller: %w", err)
	}
	fmt.Println(shops)
	var res pb.ListShopResponse
	copier.Copy(&res.Shops, shops)
	return &res, nil
}

func (r *ShopRouter) Register(server *grpc.Server) {
	pb.RegisterShopServiceServer(server, r)
}
