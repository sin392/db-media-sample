package app

import (
	"time"

	"github.com/sin392/db-media-sample/internal/infrastructure"
)

func Run() {
	app := infrastructure.NewConfig().
		Name("sample").
		ContextTimeout(10 * time.Second).
		DbNoSQL().
		WebServerPort("8080")

	app.WebServer().Start()
}
