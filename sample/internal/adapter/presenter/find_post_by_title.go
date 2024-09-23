package presenter

import (
	"github.com/sin392/db-media-sample/internal/domain/model"
	"github.com/sin392/db-media-sample/internal/usecase"
)

// var _ usecase.FindByTitlePresenter = FindByTitlePresenter{}

type FindByTitlePresenter struct{}

func NewFindByTitlePresenter() FindByTitlePresenter {
	return FindByTitlePresenter{}
}

func (p FindByTitlePresenter) Output(post *model.Post) *usecase.FindPostByTitleOutput {
	return &usecase.FindPostByTitleOutput{
		Title: post.Title,
	}
}
