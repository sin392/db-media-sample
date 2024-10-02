package controller

import (
	"net/http"

	"github.com/sin392/db-media-sample/sample/internal/adapter/controller/response"
	"github.com/sin392/db-media-sample/sample/internal/adapter/presenter"
	"github.com/sin392/db-media-sample/sample/internal/usecase"
	"github.com/sin392/db-media-sample/sample/module/trace"
)

type FindShopByNameController struct {
	uc        usecase.FindShopByNameUsecase
	presenter presenter.FindShopByNamePresenter
}

func NewFindShopByNameController(uc usecase.FindShopByNameUsecase, presenter presenter.FindShopByNamePresenter) FindShopByNameController {
	return FindShopByNameController{
		uc:        uc,
		presenter: presenter,
	}
}

func (c *FindShopByNameController) Execute(w http.ResponseWriter, r *http.Request) {
	ctx, span := trace.StartSpan(r.Context(), "FindShopByNameController.Execute")
	defer span.End()

	Name := r.URL.Query().Get("name")
	output, err := c.uc.Execute(ctx, Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	response.NewSuccess(c.presenter.Output(output), http.StatusCreated).Send(w)
}
