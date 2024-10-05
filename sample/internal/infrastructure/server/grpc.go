package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"

	"github.com/bwmarrin/snowflake"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/sin392/db-media-sample/sample/internal/adapter/controller"
	appErrors "github.com/sin392/db-media-sample/sample/internal/errors"
	pb "github.com/sin392/db-media-sample/sample/pb/shop/v1"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
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
			grpc.ChainUnaryInterceptor(
				grpc_prometheus.UnaryServerInterceptor,
				generateSnowflakeIDInterceptor(1),
				errorHandlingInterceptor,
			),
		),
	}
	server.configure(shopSrv)

	return server
}

// エラーのロギングとgRPCステータスコードの変換を行うインターセプター
func errorHandlingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	resp, err := handler(ctx, req)
	var appErr *appErrors.ApplicationError
	// エラーのロギング
	if err != nil {
		log.Printf("ERROR: %v", err)
	}
	// コンテキストに関するエラーの判定
	if errors.Is(err, context.DeadlineExceeded) {
		return nil, status.Errorf(codes.DeadlineExceeded, "timeout")
	}
	if errors.Is(err, context.Canceled) {
		return nil, status.Errorf(codes.Internal, "canceled")
	}
	// アプリケーション内部でのエラーの判定
	if errors.As(err, &appErr) {
		switch appErr.GetType() {
		case appErrors.NotFoundError:
			return nil, status.Errorf(codes.NotFound, "not found")
		case appErrors.InvalidParameterError:
			return nil, status.Errorf(codes.InvalidArgument, "invalid parameter")
		case appErrors.InternalError:
			return nil, status.Errorf(codes.Internal, "internal error")
		case appErrors.ConflictError:
			return nil, status.Errorf(codes.AlreadyExists, "conflict")
		default:
			return nil, status.Errorf(codes.Unknown, "unknown")
		}
	}
	return resp, err
}

// Snowflake ID を生成してコンテキストに追加するインターセプター
func generateSnowflakeIDInterceptor(nodeID int64) func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// Node には 1 を指定していますが、環境によって変えるべき
	node, err := snowflake.NewNode(nodeID)
	if err != nil {
		log.Fatalf("failed to create snowflake node: %v", err)
	}
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		// Generate a new Snowflake ID for each request
		snowflakeID := node.Generate().String()
		// Add the ID to the request context
		ctx = context.WithValue(ctx, "snowflakeID", snowflakeID)
		// Call the original unary handler
		return handler(ctx, req)
	}
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
