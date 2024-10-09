package server

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/sin392/db-media-sample/sample/graph"
	"google.golang.org/grpc"
)

type GqlServer struct {
	*handler.Server
}

func NewGqlServer() *GqlServer {
	return &GqlServer{}
}

func (s *GqlServer) ListenAndServe(gqlServerEndpoint string, grpcConn *grpc.ClientConn) {
	s.Server = handler.NewDefaultServer(
		graph.NewExecutableSchema(
			graph.Config{
				Resolvers: graph.NewResolver(grpcConn),
			},
		),
	)

	http.Handle("/query", s)
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.ListenAndServe(gqlServerEndpoint, nil)
}
