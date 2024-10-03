package presenter

import (
	"github.com/sin392/db-media-sample/sample/internal/domain/model"
)

type (
	FindShopByNamePresenter struct{}
	FindShopByNameOutput    struct {
		*model.Shop
	}
)

func NewFindShopByNamePresenter() FindShopByNamePresenter {
	return FindShopByNamePresenter{}
}

func (p FindShopByNamePresenter) Output(shop *model.Shop) *FindShopByNameOutput {
	return &FindShopByNameOutput{
		Shop: shop,
	}
}
