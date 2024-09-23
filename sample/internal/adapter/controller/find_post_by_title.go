package controller

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/sin392/db-media-sample/internal/adapter/presenter"
	"github.com/sin392/db-media-sample/internal/usecase"
)

type FindByTitleController struct {
	uc        usecase.FindPostByTitleUsecase
	presenter presenter.FindPostByTitlePresenter
}

func NewFindPostByTitleController(uc usecase.FindPostByTitleUsecase, presenter presenter.FindPostByTitlePresenter) FindByTitleController {
	return FindByTitleController{
		uc:        uc,
		presenter: presenter,
	}
}

func (c *FindByTitleController) Execute(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	title := r.URL.Query().Get("title")
	output, err := c.uc.Execute(ctx, title)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Header.Add(w.Header(), "content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(c.presenter.Output(output))
}
