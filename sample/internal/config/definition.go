package config

import (
	"time"
)

type Config struct {
	AppName            string        `mapstructure:"app_name"`
	AppVersion         string        `mapstructure:"app_version"`
	Timeout            time.Duration `mapstructure:"timeout"`
	GrpcServerEndpoint string        `mapstructure:"grpc_server_endpoint"`
	HttpServerEndpoint string        `mapstructure:"http_server_endpoint"`
	GqlServerEndpoint  string        `mapstructure:"gql_server_endpoint"`
}
