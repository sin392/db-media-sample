package server

import "github.com/sin392/db-media-sample/sample/internal/config"

func ExtractGrpcServerEndpointFromConfig(cfg *config.Config) (grpcServerEndpoint GrpcServerEndpoint) {
	return GrpcServerEndpoint(cfg.GrpcServerEndpoint)
}

func ExtractHttpServerEndpointFromConfig(cfg *config.Config) (httpServerEndpoint HttpServerEndpoint) {
	return HttpServerEndpoint(cfg.HttpServerEndpoint)
}

func ExtractGqlServerEndpointFromConfig(cfg *config.Config) (gqlServerEndpoint GqlServerEndpoint) {
	return GqlServerEndpoint(cfg.GqlServerEndpoint)
}
