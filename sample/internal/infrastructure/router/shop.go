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

// @Summary 名前による店舗の一致検索（単体取得）
// @Description Retrieve a shop by name
// @Accept  json
// @Produce  json
// @Param name query string true "店舗の名前"
// @Success 200 {object} usecase.FindShopByNameOutput
// @Router /shops [get]
func (r *ShopRouter) Register(router gin.IRouter) {
	shops := router.Group("/shops")
	shops.GET("", gin.WrapF(r.findShopByNameCtrl.Execute))
}
