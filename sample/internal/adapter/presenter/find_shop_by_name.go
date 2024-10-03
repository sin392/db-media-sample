package presenter

import (
	"github.com/sin392/db-media-sample/sample/internal/domain/model"
)

type (
	FindShopByNamePresenter struct{}
	FindShopByNameOutput    struct {
		model.Shop
	}
)

func NewFindShopByNamePresenter() FindShopByNamePresenter {
	return FindShopByNamePresenter{}
}

func (p FindShopByNamePresenter) Output(shop *model.Shop) *FindShopByNameOutput {
	return &FindShopByNameOutput{
		Shop: model.Shop{
			ID:       shop.ID,
			Name:     shop.Name,
			Location: shop.Location,
			Tel:      shop.Tel,
			ImageURL: shop.ImageURL,
			SiteURL:  shop.SiteURL,
			Rating:   shop.Rating,
			Tags:     shop.Tags,
			Menus:    shop.Menus,
		},
	}
}
