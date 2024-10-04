package presenter

import (
	"fmt"

	"github.com/jinzhu/copier"
	"github.com/sin392/db-media-sample/sample/internal/usecase"
)

// ViewModel
type (
	ShopResponse struct {
		usecase.ShopOutput
	}
	ShopListResponse struct {
		usecase.ShopListOutput
	}
)

// ShopOutputをShopListResponseに変換する
func OutputShopPassThrough(output *usecase.ShopOutput) (*ShopResponse, error) {
	var res ShopResponse
	if err := copier.Copy(&res, output); err != nil {
		return nil, fmt.Errorf("failed to copy output to res: %w", err)
	}
	return &res, nil
}

// ShopListOutputをShopListResponseに変換する
func OutputShopListPassThrough(output *usecase.ShopListOutput) (*ShopListResponse, error) {
	var res ShopListResponse
	if err := copier.Copy(&res, output); err != nil {
		return nil, fmt.Errorf("failed to copy output to res: %w", err)
	}
	return &res, nil
}
