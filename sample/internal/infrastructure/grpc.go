package infrastructure

import "github.com/sin392/db-media-sample/sample/pb/shop/v1"

type GrpcServer struct {
	shop.UnimplementedShopServiceServer
}
