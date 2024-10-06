package server

import (
	"fmt"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/sin392/db-media-sample/sample/graph"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GqlServer struct {
	*handler.Server
}

func NewGqlServer() GqlServer {
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
		panic(fmt.Errorf("failed to connect to grpc server: %w", err))
	}
	srv := handler.NewDefaultServer(
		graph.NewExecutableSchema(
			graph.Config{
				Resolvers: graph.NewResolver(conn),
			},
		),
	)
	return GqlServer{
		Server: srv,
	}
}

func (s *GqlServer) ListenAndServe() {
	http.Handle("/query", s)
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.ListenAndServe(":8081", nil)
}
