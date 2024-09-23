// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package app

import (
	"github.com/google/wire"
	"github.com/sin392/db-media-sample/internal/adapter/controller"
	"github.com/sin392/db-media-sample/internal/adapter/presenter"
	"github.com/sin392/db-media-sample/internal/adapter/repositoryimpl/nosql"
	"github.com/sin392/db-media-sample/internal/infrastructure/database"
	"github.com/sin392/db-media-sample/internal/infrastructure/router"
	"github.com/sin392/db-media-sample/internal/usecase"
)

// Injectors from wire.go:

func InitializeApplication() (*Application, error) {
	config := database.NewConfig()
	noSQL, err := database.NewMongoHandler(config)
	if err != nil {
		return nil, err
	}
	shopRepository := nosql.NewShopRepositoryImpl(noSQL)
	findShopByNameUsecase := usecase.NewFindShopByNameIntercepter(shopRepository)
	findShopByNamePresenter := presenter.NewFindShopByNamePresenter()
	findShopByNameController := controller.NewFindShopByNameController(findShopByNameUsecase, findShopByNamePresenter)
	shopRouter := router.NewShopRouter(findShopByNameController)
	routers := router.NewRouters(shopRouter)
	application, err := NewApplication(routers)
	if err != nil {
		return nil, err
	}
	return application, nil
}

// wire.go:

var WireSet = wire.NewSet(router.NewRouters, router.NewShopRouter, controller.NewFindShopByNameController, presenter.NewFindShopByNamePresenter, usecase.NewFindShopByNameIntercepter, nosql.NewShopRepositoryImpl, database.NewMongoHandler, database.NewConfig)
