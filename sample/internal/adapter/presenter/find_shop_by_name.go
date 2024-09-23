package presenter

import (
	"github.com/sin392/db-media-sample/internal/domain/model"
	"github.com/sin392/db-media-sample/internal/usecase"
)

type FindShopByNamePresenter struct{}

func NewFindShopByNamePresenter() FindShopByNamePresenter {
	return FindShopByNamePresenter{}
}

func (p FindShopByNamePresenter) Output(Shop *model.Shop) *usecase.FindShopByNameOutput {
	return &usecase.FindShopByNameOutput{
		Name:   Shop.Name,
		Tel:    Shop.Tel,
		Rating: Shop.Rating,
	}
}
