package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sin392/db-media-sample/sample/pb/shop/v1"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
)

type HttpServer struct {
	mux http.Handler
}

func NewHttpServer() *HttpServer {
	return &HttpServer{}
}

func (s *HttpServer) ListenAndServe(httpServerEndpoint string, grpcConn *grpc.ClientConn) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := s.setupRouters(ctx, grpcConn); err != nil {
		panic(err)
	}

	return http.ListenAndServe(httpServerEndpoint, s.mux)
}

func (s *HttpServer) setupRouters(ctx context.Context, grpcConn *grpc.ClientConn) error {
	mux := runtime.NewServeMux(
		runtime.WithHealthzEndpoint(
			grpc_health_v1.NewHealthClient(grpcConn),
		),
	)
	if err := shop.RegisterShopServiceHandler(ctx, mux, grpcConn); err != nil {
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
