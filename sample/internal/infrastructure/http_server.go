package infrastructure

import (
	"time"

	"github.com/sin392/db-media-sample/internal/infrastructure/router"
)

type Config struct {
	appName       string
	ctxTimeout    time.Duration
	webServerPort router.Port
}

func NewConfig(
	appName string,
	webServerPort router.Port,
	ctxTimeout time.Duration,
) *Config {
	return &Config{
		appName:       appName,
		ctxTimeout:    ctxTimeout,
		webServerPort: webServerPort,
	}
}

// configとサーバ関連のコード配置整備したい
func NewServer(cfg *Config, routers router.Routers) router.Server {
	return router.NewGinServer(
		cfg.webServerPort,
		cfg.ctxTimeout,
		routers,
	)
}
