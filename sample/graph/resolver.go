package graph

import (
	pb "github.com/sin392/db-media-sample/sample/pb/shop/v1"
	"google.golang.org/grpc"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	srv pb.ShopServiceClient
}

func NewResolver(conn *grpc.ClientConn) *Resolver {
	srv := pb.NewShopServiceClient(conn)
	return &Resolver{
		srv: srv,
	}
}
