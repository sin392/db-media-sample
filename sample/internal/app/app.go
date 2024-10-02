package app

import (
	"github.com/sin392/db-media-sample/sample/internal/config"
	"github.com/sin392/db-media-sample/sample/internal/infrastructure"
	"github.com/sin392/db-media-sample/sample/module/trace"
)

type Application struct {
	cfg    *config.Config
	server infrastructure.Server
}

func NewApplication(
	cfg *config.Config,
	server infrastructure.Server,
) (*Application, error) {
	return &Application{
		cfg:    cfg,
		server: server,
	}, nil
}

func (a *Application) Configure() {
	// OpenTelemetryの初期化
	trace.InitTraceProvider(a.cfg)
}

func (a *Application) Start() {
	a.server.ListenAndServe()
}
