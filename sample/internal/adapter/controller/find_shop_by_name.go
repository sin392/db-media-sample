package controller

import (
	"net/http"

	"github.com/sin392/db-media-sample/internal/adapter/controller/response"
	"github.com/sin392/db-media-sample/internal/adapter/presenter"
	"github.com/sin392/db-media-sample/internal/usecase"
	"github.com/sin392/db-media-sample/module/trace"
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
	ctx, span := trace.StartSpan(r.Context(), "FindByNameController.Execute")
	defer span.End()

	Name := r.URL.Query().Get("name")
	output, err := c.uc.Execute(ctx, Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	response.NewSuccess(c.presenter.Output(output), http.StatusCreated).Send(w)
}
