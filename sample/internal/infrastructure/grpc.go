package infrastructure

import (
	"fmt"
	"log"
	"net"

	"github.com/sin392/db-media-sample/sample/internal/infrastructure/router"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

type GrpcServer struct {
	server     *grpc.Server
	shopRouter router.ShopRouter
}

func NewGrpcServer(
	shopRouter router.ShopRouter,
) GrpcServer {
	server := GrpcServer{
		server: grpc.NewServer(
			grpc.StatsHandler(
				otelgrpc.NewServerHandler(),
			),
		),
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

func (s *GrpcServer) setupRouters(routers GrpcRouters) {
	hs := health.NewServer()
	grpc_health_v1.RegisterHealthServer(s.server, hs)
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
