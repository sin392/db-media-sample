package presenter

import (
	"github.com/sin392/db-media-sample/sample/internal/domain/model"
)

type (
	ListShopPresenter struct{}
	ListShopOutput    struct {
		shops []*model.Shop
	}
)

func NewListShopPresenter() ListShopPresenter {
	return ListShopPresenter{}
}

func (p ListShopPresenter) Output(shops []*model.Shop) *ListShopOutput {
	return &ListShopOutput{
		shops: shops,
	}
}
