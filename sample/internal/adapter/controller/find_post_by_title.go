package controller

import (
	"net/http"

	"github.com/sin392/db-media-sample/internal/adapter/controller/response"
	"github.com/sin392/db-media-sample/internal/adapter/presenter"
	"github.com/sin392/db-media-sample/internal/usecase"
)

type FindByTitleController struct {
	uc        usecase.FindPostByTitleUsecase
	presenter presenter.FindByTitlePresenter
}

func NewFindByTitleController(uc usecase.FindPostByTitleUsecase, presenter presenter.FindByTitlePresenter) FindByTitleController {
	return FindByTitleController{
		uc:        uc,
		presenter: presenter,
	}
}

func (c *FindByTitleController) Execute(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	title := r.URL.Query().Get("title")
	output, err := c.uc.Execute(ctx, title)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	response.NewSuccess(c.presenter.Output(output), http.StatusCreated).Send(w)
}
