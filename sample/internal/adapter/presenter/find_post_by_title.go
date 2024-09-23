package presenter

import (
	"github.com/sin392/db-media-sample/internal/domain/model"
	"github.com/sin392/db-media-sample/internal/usecase"
)

// var _ usecase.FindPostByTitlePresenter = findPostByTitlePresenter{}

type FindPostByTitlePresenter struct{}

func NewFindPostByTitlePresenter() FindPostByTitlePresenter {
	return FindPostByTitlePresenter{}
}

func (p FindPostByTitlePresenter) Output(post *model.Post) *usecase.FindPostByTitleOutput {
	return &usecase.FindPostByTitleOutput{
		Title: post.Title,
	}
}
