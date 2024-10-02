package infrastructure

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/sin392/db-media-sample/sample/internal/infrastructure/router"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

type GrpcServer struct {
	server     *grpc.Server
	shopRouter router.ShopRouter
}

func NewGrpcServer(
	shopRouter router.ShopRouter,
) GrpcServer {
	server := GrpcServer{
		server:     grpc.NewServer(),
		shopRouter: shopRouter,
	}
	reflection.Register(server.server)
	routers := NewGrpcRouters(shopRouter)
	server.setupRouters(routers)

	return server
}

type GrpcRouter interface {
	Register(server *grpc.Server)
}

type GrpcRouters []GrpcRouter

func NewGrpcRouters(
	shopRouter router.ShopRouter,
) GrpcRouters {
	return GrpcRouters{
		&shopRouter,
	}
}

func (s *GrpcServer) Check(ctx context.Context, req *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	return &grpc_health_v1.HealthCheckResponse{
		Status: grpc_health_v1.HealthCheckResponse_SERVING,
	}, nil
}

func (s *GrpcServer) Watch(req *grpc_health_v1.HealthCheckRequest, w grpc_health_v1.Health_WatchServer) error {
	return status.Error(codes.Unimplemented, "Watching is not supported")
}

func (s *GrpcServer) setupRouters(routers GrpcRouters) {
	grpc_health_v1.RegisterHealthServer(s.server, s)
	for _, router := range routers {
		router.Register(s.server)
	}
}

func (s *GrpcServer) ListenAndServe() error {
	listenPort, err := net.Listen("tcp", ":50051")
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}
	if err := s.server.Serve(listenPort); err != nil {
		return fmt.Errorf("failed to serve: %v", err)
	}
	log.Println("gRPC server started")
	return nil
}
