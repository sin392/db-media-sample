package app

import (
	"time"

	"github.com/sin392/db-media-sample/internal/infrastructure"
	"github.com/sin392/db-media-sample/module/trace"
)

type Application struct {
	routers infrastructure.Routers
	server  infrastructure.Server
}

func NewApplication(
	routers infrastructure.Routers,
) (*Application, error) {
	return &Application{
		routers: routers,
		server:  nil,
	}, nil
}

func (a *Application) Configure() {
	// OpenTelemetryの初期化
	trace.InitTraceProvider()
	// 設定値の読み込み
	config := infrastructure.NewConfig("sample", infrastructure.Port(8080), 10*time.Second)

	// サーバーの初期化
	a.server = infrastructure.NewServer(config, a.routers)
}

func (a *Application) Start() {
	a.server.ListenAndServe()
}
