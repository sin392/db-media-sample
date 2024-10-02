package infrastructure

import (
	"context"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/sin392/db-media-sample/sample/pb/shop/v1"
	"google.golang.org/grpc"
)

type HttpServer struct {
	mux *runtime.ServeMux
}

func NewHttpServer() HttpServer {
	server := HttpServer{}
	return server
}

func (s *HttpServer) setupRouters(ctx context.Context) error {
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	grpcServerEndpoint := "localhost:50051"
	err := shop.RegisterShopServiceHandlerFromEndpoint(ctx, mux, grpcServerEndpoint, opts)
	if err != nil {
		return err
	}
	s.mux = mux
	return nil
}

func (s *HttpServer) ListenAndServe() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := s.setupRouters(ctx); err != nil {
		panic(err)
	}
	return http.ListenAndServe(":8080", s.mux)
}
