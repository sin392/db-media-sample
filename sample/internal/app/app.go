package app

import (
	"time"

	"github.com/sin392/db-media-sample/internal/infrastructure"
	"github.com/sin392/db-media-sample/module/trace"
)

type Application struct {
	config *infrastructure.Config
}

func NewApplication(
	config *infrastructure.Config,
) (*Application, error) {
	return &Application{
		config: config,
	}, nil
}

func (a *Application) Configure() {
	// OpenTelemetryの初期化
	trace.InitTraceProvider()
	// 設定値の読み込み
	a.config.
		Name("sample").
		ContextTimeout(10 * time.Second).
		DbNoSQL().
		WebServerPort("8080")
}

func (a *Application) Start() {
	a.config.WebServer().Start()
}
