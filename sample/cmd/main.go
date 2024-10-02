package main

import "github.com/sin392/db-media-sample/sample/internal/app"

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @BasePath /v1
func main() {
	app, err := app.InitializeApplication()
	if err != nil {
		panic(err)
	}

	app.Configure()

	app.Start()
}
