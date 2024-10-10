package server

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/sin392/db-media-sample/sample/graph"
	"google.golang.org/grpc"
)

type GqlServerEndpoint string

func (e GqlServerEndpoint) String() string {
	return string(e)
}

type GqlServer struct {
	*handler.Server
	gqlServerEndpoint GqlServerEndpoint
}

func NewGqlServer(gqlServerEndpoint GqlServerEndpoint, grpcConn *grpc.ClientConn) *GqlServer {
	server := &GqlServer{
		Server: handler.NewDefaultServer(
			graph.NewExecutableSchema(
				graph.Config{
					Resolvers: graph.NewResolver(grpcConn),
				},
			),
		),
		gqlServerEndpoint: gqlServerEndpoint,
	}

	http.Handle("/query", server)
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))

	return server
}

func (s *GqlServer) ListenAndServe() {
	http.ListenAndServe(s.gqlServerEndpoint.String(), nil)
}
