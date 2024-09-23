package controller

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/sin392/db-media-sample/internal/adapter/presenter"
	"github.com/sin392/db-media-sample/internal/usecase"
)

type FindByNameController struct {
	uc        usecase.FindShopByNameUsecase
	presenter presenter.FindShopByNamePresenter
}

func NewFindShopByNameController(uc usecase.FindShopByNameUsecase, presenter presenter.FindShopByNamePresenter) FindByNameController {
	return FindByNameController{
		uc:        uc,
		presenter: presenter,
	}
}

func (c *FindByNameController) Execute(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	Name := r.URL.Query().Get("name")
	output, err := c.uc.Execute(ctx, Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Header.Add(w.Header(), "content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(c.presenter.Output(output))
}
