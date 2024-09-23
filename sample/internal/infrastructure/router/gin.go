package router

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sin392/db-media-sample/internal/adapter/controller"
	"github.com/sin392/db-media-sample/internal/adapter/presenter"
	"github.com/sin392/db-media-sample/internal/adapter/repositoryimpl/nosql"
	"github.com/sin392/db-media-sample/internal/usecase"
)

type Server interface {
	Listen()
}

type Port int64

type ginEngine struct {
	router     *gin.Engine
	db         nosql.NoSQL
	port       Port
	ctxTimeout time.Duration
}

func NewGinServer(
	db nosql.NoSQL,
	port Port,
	t time.Duration,
) *ginEngine {
	return &ginEngine{
		router:     gin.New(),
		db:         db,
		port:       port,
		ctxTimeout: t,
	}
}

func (g ginEngine) Listen() {
	gin.SetMode(gin.ReleaseMode)
	gin.Recovery()

	g.setAppHandlers(g.router)

	server := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 15 * time.Second,
		Addr:         fmt.Sprintf(":%d", g.port),
		Handler:      g.router,
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := server.ListenAndServe(); err != nil {
			fmt.Println(err)
		}
	}()

	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()

	if err := server.Shutdown(ctx); err != nil {
		fmt.Println(err)
	}
}

func (g ginEngine) setAppHandlers(router *gin.Engine) {
	router.GET("/v1/health", g.healthcheck())
	router.GET("/v1/posts", g.buildNewFindByTitleAction())
	router.GET("/v1/shops", g.buildNewFindShopByNameAction())
}

func (g ginEngine) buildNewFindByTitleAction() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			uc = usecase.NewFindPostByTitleIntercepter(
				nosql.NewPostRepositoryImpl(g.db),
				g.ctxTimeout,
			)
			ctrl = controller.NewFindByTitleController(uc, presenter.NewFindByTitlePresenter())
		)

		ctrl.Execute(c.Writer, c.Request)
	}
}

func (g ginEngine) buildNewFindShopByNameAction() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			uc = usecase.NewFindShopByNameIntercepter(
				nosql.NewShopRepositoryImpl(g.db),
				g.ctxTimeout,
			)
			ctrl = controller.NewFindShopByNameController(uc, presenter.NewFindShopByNamePresenter())
		)

		ctrl.Execute(c.Writer, c.Request)
	}
}

func (g ginEngine) healthcheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	}
}
