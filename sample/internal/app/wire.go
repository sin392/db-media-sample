//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"

	"github.com/sin392/db-media-sample/sample/internal/adapter/controller"
	"github.com/sin392/db-media-sample/sample/internal/adapter/presenter"
	"github.com/sin392/db-media-sample/sample/internal/adapter/repositoryimpl/nosql"
	"github.com/sin392/db-media-sample/sample/internal/config"
	"github.com/sin392/db-media-sample/sample/internal/infrastructure"
	"github.com/sin392/db-media-sample/sample/internal/infrastructure/database"
	"github.com/sin392/db-media-sample/sample/internal/infrastructure/router"
	"github.com/sin392/db-media-sample/sample/internal/usecase"
)

var WireSet = wire.NewSet(
	config.Load,
	// infrastructure
	infrastructure.NewHttpServer,
	infrastructure.NewGrpcServer,
	router.NewShopRouter,
	database.NewMongoHandler,
	database.NewConfig,
	// adapter
	controller.NewFindShopByNameController,
	presenter.NewFindShopByNamePresenter,
	// usecase
	usecase.NewFindShopByNameIntercepter,
	nosql.NewShopRepositoryImpl,
)

func InitializeApplication() (*Application, error) {
	wire.Build(
		WireSet,
		NewApplication,
	)
	return nil, nil
}
