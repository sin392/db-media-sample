package router

import (
	"context"

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
	}
}

func (r *ShopRouter) FindShopByName(ctx context.Context, req *pb.FindShopByNameRequest) (*pb.FindShopByNameResponse, error) {
	name := req.GetName()
	shop, err := r.findShopByNameCtrl.Execute(ctx, name)
	if err != nil {
		return nil, err
	}
	menus := make([]*pb.Menu, len(shop.Menus))
	for i, m := range shop.Menus {
		menus[i] = &pb.Menu{
			Name:  m.Name,
			Price: int32(m.Price),
			Desc:  m.Desc,
		}
	}
	return &pb.FindShopByNameResponse{
		Id:   shop.ID,
		Name: shop.Name,
		Location: &pb.Location{
			Prefecture: shop.Location.Prefecture,
			City:       shop.Location.City,
			Address:    shop.Location.Address,
		},
		Tel:      shop.Tel,
		ImageUrl: shop.ImageURL,
		SiteUrl:  shop.SiteURL,
		Rating:   shop.Rating,
		Tags:     shop.Tags,
		Menus:    menus,
	}, nil
}

func (r *ShopRouter) ListShop(ctx context.Context, req *pb.ListShopRequest) (*pb.ListShopResponse, error) {
	_, err := r.listShopCtrl.Execute(ctx)
	if err != nil {
		return nil, err
	}
	return &pb.ListShopResponse{
		Shops: []*pb.Shop{},
	}, nil
}

func (r *ShopRouter) Register(server *grpc.Server) {
	pb.RegisterShopServiceServer(server, r)
}
