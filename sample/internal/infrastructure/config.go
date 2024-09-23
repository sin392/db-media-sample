package infrastructure

import (
	"time"
)

type Port int64

type Config struct {
	AppName       string
	CtxTimeout    time.Duration
	WebServerPort Port
}

func NewConfig(
	appName string,
	webServerPort Port,
	ctxTimeout time.Duration,
) *Config {
	return &Config{
		AppName:       appName,
		CtxTimeout:    ctxTimeout,
		WebServerPort: webServerPort,
	}
}
