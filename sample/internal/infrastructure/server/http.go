package server

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sin392/db-media-sample/sample/pb/shop/v1"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
)

type HttpServerEndpoint string

func (e HttpServerEndpoint) String() string {
	return string(e)
}

type HttpServer struct {
	*runtime.ServeMux
	httpServerEndpoint HttpServerEndpoint
}

func NewHttpServer(httpServerEndpoint HttpServerEndpoint, grpcConn *grpc.ClientConn) (*HttpServer, error) {
	server := &HttpServer{
		ServeMux: runtime.NewServeMux(
			runtime.WithHealthzEndpoint(
				grpc_health_v1.NewHealthClient(grpcConn),
			),
		),
		httpServerEndpoint: httpServerEndpoint,
	}
	ctx := context.Background()
	// rpcサービスのエンドポイント
	if err := shop.RegisterShopServiceHandler(ctx, server.ServeMux, grpcConn); err != nil {
		return nil, fmt.Errorf("failed to register shop service handler: %w", err)
	}
	// メトリクスエンドポイント
	if err := server.HandlePath("GET", "/metrics",
		runtime.HandlerFunc(func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
			// TODO: otel形式でエクスポートするようにする
			metricHandler := promhttp.Handler()
			metricHandler.ServeHTTP(w, r)
		}),
	); err != nil {
		return nil, fmt.Errorf("failed to handle metrics: %w", err)
	}
	// Swaggerエンドポイント
	// TODO: 開発環境以外では公開しないようにする
	if err := server.HandlePath("GET", "/docs/swagger.yaml", func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		http.ServeFile(w, r, "./docs/openapiv2/apidocs.swagger.yaml")
	}); err != nil {
		return nil, fmt.Errorf("failed to handle swagger.yaml: %w", err)
	}
	// SwaggerUIエンドポイント
	if err := server.HandlePath("GET", "/docs",
		runtime.HandlerFunc(func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
			swaggerHandler := middleware.SwaggerUI(middleware.SwaggerUIOpts{
				SpecURL: "/docs/swagger.yaml",
			}, server)
			swaggerHandler.ServeHTTP(w, r)
		}),
	); err != nil {
		return nil, fmt.Errorf("failed to handle swagger ui: %w", err)
	}

	return server, nil
}

// リクエストのメソッドとパスからスパンの名称を構成するミドルウェア.
func traceMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		operation := fmt.Sprintf("%s %s", r.Method, r.URL.Path)
		handler := otelhttp.NewHandler(next, operation)
		handler.ServeHTTP(w, r)
	})
}

// ポートの指定とコネクションの指定がずれてるの微妙だな
// ポートの設定もファクトリ側に持ってくべきか.
func (s *HttpServer) ListenAndServe() error {
	return http.ListenAndServe(s.httpServerEndpoint.String(), traceMiddleware(s))
}

func (s *HttpServer) Shutdown(ctx context.Context) error {
	if err := s.Shutdown(ctx); err != nil {
		return fmt.Errorf("failed to shutdown http server: %w", err)
	}

	log.Println("http server stopped")

	return nil
}
