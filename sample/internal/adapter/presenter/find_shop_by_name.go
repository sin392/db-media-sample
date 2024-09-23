package presenter

import (
	"github.com/sin392/db-media-sample/internal/domain/model"
	"github.com/sin392/db-media-sample/internal/usecase"
)

type FindShopByNamePresenter struct{}

func NewFindShopByNamePresenter() FindShopByNamePresenter {
	return FindShopByNamePresenter{}
}

func (p FindShopByNamePresenter) Output(shop *model.Shop) *usecase.FindShopByNameOutput {
	return &usecase.FindShopByNameOutput{
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
