package router

import (
	"github.com/gin-gonic/gin"
	"github.com/sin392/db-media-sample/internal/adapter/controller"
)

type ShopRouter struct {
	findShopByNameCtrl controller.FindShopByNameController
}

func NewShopRouter(
	findShopByNameCtrl controller.FindShopByNameController,
) ShopRouter {
	return ShopRouter{
		findShopByNameCtrl: findShopByNameCtrl,
	}
}

func (r *ShopRouter) Register(router gin.IRouter) {
	shops := router.Group("/shops")
	shops.GET("", gin.WrapF(r.findShopByNameCtrl.Execute))
}
