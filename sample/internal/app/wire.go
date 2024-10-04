//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"

	"github.com/sin392/db-media-sample/sample/internal/adapter/controller"
	"github.com/sin392/db-media-sample/sample/internal/adapter/repository"
	"github.com/sin392/db-media-sample/sample/internal/config"
	"github.com/sin392/db-media-sample/sample/internal/infrastructure/database"
	"github.com/sin392/db-media-sample/sample/internal/infrastructure/server"
	"github.com/sin392/db-media-sample/sample/internal/usecase"
)

// usecaseとadapterが増えるとファイルが肥大化しそうだなぁ
// presenterは汎用的な表現にまとめていいかも
var WireSet = wire.NewSet(
	config.Load,
	// infrastructure
	server.NewHttpServer,
	server.NewGrpcServer,
	database.NewMongoHandler,
	database.NewConfig,
	// adapter
	controller.NewShopPbController,
	repository.NewShopNoSQLQueryRepository,
	// usecase
	usecase.NewFindShopByNameUsecase,
	usecase.NewListShopUsecase,
)

func InitializeApplication() (*Application, error) {
	wire.Build(
		WireSet,
		NewApplication,
	)
	return nil, nil
}
