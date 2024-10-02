package infrastructure

import (
	"context"
	"fmt"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/sin392/db-media-sample/sample/pb/shop/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type HttpServer struct {
	mux *runtime.ServeMux
}

func NewHttpServer() HttpServer {
	server := HttpServer{}
	return server
}

func (s *HttpServer) setupRouters(ctx context.Context) error {
	grpcServerEndpoint := "localhost:50051"
	conn, err := grpc.NewClient(
		grpcServerEndpoint,
		grpc.WithTransportCredentials(
			insecure.NewCredentials(),
		),
	)
	if err != nil {
		return fmt.Errorf("failed to connect to grpc server: %w", err)
	}
	mux := runtime.NewServeMux()
	s.mux = mux
	if err := shop.RegisterShopServiceHandler(ctx, s.mux, conn); err != nil {
		return err
	}
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
