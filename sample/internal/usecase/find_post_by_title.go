package usecase

import (
	"context"
	"time"

	"github.com/sin392/db-media-sample/internal/domain/model"
	"github.com/sin392/db-media-sample/internal/domain/repository"
)

type FindPostByTitleUsecase interface {
	Execute(ctx context.Context, title string) (*model.Post, error)
}

// PresenterをControllerで呼ぶ場合はDIP不要
// type FindPostByTitlePresenter interface {
// 	Output(post *model.Post) *FindPostByTitleOutput
// }

type FindPostByTitleOutput struct {
	Title string `json:"title"`
}

type findPostByTitleInteractor struct {
	repo       repository.PostRepository
	ctxTimeout time.Duration
}

func NewFindPostByTitleIntercepter(
	repo repository.PostRepository,
	ctxTimeout time.Duration,
) FindPostByTitleUsecase {
	return &findPostByTitleInteractor{
		repo:       repo,
		ctxTimeout: ctxTimeout,
	}
}

func (a *findPostByTitleInteractor) Execute(ctx context.Context, title string) (*model.Post, error) {
	post, err := a.repo.FindByTitle(ctx, title)
	if err != nil {
		return nil, err
	}
	return post, nil
}
