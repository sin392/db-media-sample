package infrastructure

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sin392/db-media-sample/sample/pb/shop/v1"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/health/grpc_health_v1"
)

type HttpServer struct {
	mux http.Handler
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
		grpc.WithStatsHandler(
			otelgrpc.NewClientHandler(),
		),
	)
	if err != nil {
		return fmt.Errorf("failed to connect to grpc server: %w", err)
	}
	mux := runtime.NewServeMux(
		runtime.WithHealthzEndpoint(
			grpc_health_v1.NewHealthClient(conn),
		),
	)
	if err := shop.RegisterShopServiceHandler(ctx, mux, conn); err != nil {
		return err
	}
	// メトリクスエンドポイント
	mux.HandlePath("GET", "/metrics",
		runtime.HandlerFunc(func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
			metricHandler := promhttp.Handler()
			metricHandler.ServeHTTP(w, r)
		}),
	)
	// Swaggerエンドポイント
	// TODO: 開発環境以外では公開しないようにする
	mux.HandlePath("GET", "/docs/swagger.yaml", func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		http.ServeFile(w, r, "./docs/openapiv2/apidocs.swagger.yaml")
	})
	// SwaggerUIエンドポイント
	mux.HandlePath("GET", "/docs",
		runtime.HandlerFunc(func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
			swaggerHandler := middleware.SwaggerUI(middleware.SwaggerUIOpts{
				SpecURL: "/docs/swagger.yaml",
			}, mux)
			swaggerHandler.ServeHTTP(w, r)
		}),
	)
	// 以下はmiddleware的に設定できるのでは？
	s.mux = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// リクエストのパスをスパンの名前とするHTTPハンドラを生成
		otelHandler := otelhttp.NewHandler(mux, fmt.Sprintf("%s %s", r.Method, r.URL.Path))
		otelHandler.ServeHTTP(w, r)
	})
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
