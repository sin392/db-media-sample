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

// usecaseとadapterが増えるとファイルが肥大化しそうだなぁ
// presenterは汎用的な表現にまとめていいかも
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
	controller.NewListShopController,
	presenter.NewFindShopByNamePresenter,
	presenter.NewListShopPresenter,
	// usecase
	usecase.NewFindShopByNameIntercepter,
	usecase.NewListShopIntercepter,
	nosql.NewShopRepositoryImpl,
)

func InitializeApplication() (*Application, error) {
	wire.Build(
		WireSet,
		NewApplication,
	)
	return nil, nil
}
