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
	res, err := r.findShopByNameCtrl.Execute(ctx, name)
	if err != nil {
		return nil, fmt.Errorf("failed to execute controller: %w", err)
	}
	var pbRes pb.FindShopByNameResponse
	if err := copier.Copy(&pbRes, res); err != nil {
		return nil, fmt.Errorf("failed to copy from res to pbRes: %w", err)
	}
	return &pbRes, nil
}

func (r *ShopRouter) ListShop(ctx context.Context, req *pb.ListShopRequest) (*pb.ListShopResponse, error) {
	res, err := r.listShopCtrl.Execute(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to execute controller: %w", err)
	}
	var pbRes pb.ListShopResponse
	// ここのコピーもう少し改善できないか？
	if err := copier.Copy(&pbRes.Shops, res.ShopListOutput.ShopList); err != nil {
		return nil, fmt.Errorf("failed to copy from res to pbRes: %w", err)
	}
	return &pbRes, nil
}

func (r *ShopRouter) Register(server *grpc.Server) {
	pb.RegisterShopServiceServer(server, r)
}
