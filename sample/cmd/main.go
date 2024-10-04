package main

import "github.com/sin392/db-media-sample/sample/internal/app"

func main() {
	app, err := app.InitializeApplication()
	if err != nil {
		panic(err)
	}

	app.Start()
}
