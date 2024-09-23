package app

import (
	"time"

	"github.com/sin392/db-media-sample/internal/infrastructure"
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
	a.config.
		Name("sample").
		ContextTimeout(10 * time.Second).
		DbNoSQL().
		WebServerPort("8080")
}

func (a *Application) Start() {
	a.config.WebServer().Start()
}
