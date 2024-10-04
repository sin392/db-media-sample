package server

import (
	"fmt"
	"log"
	"net"

	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/sin392/db-media-sample/sample/internal/adapter/controller"
	pb "github.com/sin392/db-media-sample/sample/pb/shop/v1"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

type GrpcServer struct {
	*grpc.Server
}

func NewGrpcServer(
	shopSrv controller.ShopControllerPb,
) GrpcServer {
	server := GrpcServer{
		Server: grpc.NewServer(
			grpc.StatsHandler(
				otelgrpc.NewServerHandler(),
			),
			grpc.StreamInterceptor(grpc_prometheus.StreamServerInterceptor),
			grpc.UnaryInterceptor(grpc_prometheus.UnaryServerInterceptor),
		),
	}
	server.configure(shopSrv)

	return server
}

func (s *GrpcServer) configure(shopSrv controller.ShopControllerPb) {
	// リフレクションサービスの登録
	reflection.Register(s)
	// ヘルスチェックサービスの登録
	grpc_health_v1.RegisterHealthServer(s, health.NewServer())
	// gRPCのメトリクスサービスの登録
	grpc_prometheus.Register(s.Server)
	// Shopサービスの登録
	pb.RegisterShopServiceServer(s, &shopSrv)
}

func (s *GrpcServer) ListenAndServe() error {
	listenPort, err := net.Listen("tcp", ":50051")
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}
	if err := s.Serve(listenPort); err != nil {
		return fmt.Errorf("failed to serve: %v", err)
	}
	log.Println("gRPC server started")
	return nil
}
