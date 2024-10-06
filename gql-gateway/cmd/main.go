// gateway/main.go
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	pb "github.com/sin392/db-media-sample/sample/pb/shop/v1"
	"github.com/ysugimoto/grpc-graphql-gateway/runtime"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// 別コンテナへのアクセスなのでホストにはサービス名を指定
	grpcServerEndpoint := "api:50051"
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
		log.Fatalln(fmt.Errorf("failed to connect to grpc server: %w", err))
	}
	// マルチプレクサはgRPCへのリクエストの振り分けを行う
	mux := runtime.NewServeMux(
	// ヘルスチェックの登録どうやったら？
	// runtime.MiddlewareFunc(
	// 	grpc_health_v1.NewHealthClient(conn),
	// ),
	)
	c := graphql.SchemaConfig{}
	schema, _ := graphql.NewSchema(c)
	h := handler.New(&handler.Config{
		Schema:     &schema,
		Pretty:     true,
		GraphiQL:   false,
		Playground: true,
		PlaygroundConfig: &handler.PlaygroundConfig{
			Endpoint:             "/query",
			SubscriptionEndpoint: "/query",
		},
	})

	if err := pb.RegisterShopServiceGraphqlHandler(mux, conn); err != nil {
		log.Fatalln(err)
	}
	// playgroundのエンドポイント
	http.Handle("/playground", h)
	// graphqlのエンドポイント
	http.Handle("/query", mux)
	log.Fatalln(http.ListenAndServe(":8082", nil))
}
