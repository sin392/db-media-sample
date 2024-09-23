package infrastructure

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sin392/db-media-sample/internal/infrastructure/router"
	ginprometheus "github.com/zsais/go-gin-prometheus"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

type Server interface {
	ListenAndServe()
}

type Router interface {
	Register(router gin.IRouter)
}

type Routers []Router

// TODO: フィールド整備
type ginEngine struct {
	router  *gin.Engine
	routers Routers
	server  *http.Server
}

func NewRouters(
	shopRouter router.ShopRouter,
) Routers {
	return Routers{
		&shopRouter,
	}
}

func NewServer(
	cfg *Config,
	routers Routers,
) Server {
	router := gin.Default()

	server := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 15 * time.Second,
		Addr:         fmt.Sprintf(":%d", cfg.WebServerPort),
		Handler:      router,
	}

	gh := ginEngine{
		router:  router,
		routers: routers,
		server:  server,
	}
	gh.useMiddleWares()

	return &gh
}

func (g *ginEngine) useMiddleWares() {
	// openteremetryの設定
	// TODO: どこかに切り出したい
	g.router.ContextWithFallback = true
	g.router.Use(otelgin.Middleware("db-media-sample"))
	// metric exporterの設定
	p := ginprometheus.NewPrometheus("gin")
	p.ReqCntURLLabelMappingFn = func(c *gin.Context) string {
		url := c.Request.URL.Path
		for _, p := range c.Params {
			if p.Key == "name" {
				url = strings.Replace(url, p.Value, ":name", 1)
				break
			}
		}
		return url
	}
	p.Use(g.router)
}

func (g *ginEngine) ListenAndServe() {
	gin.SetMode(gin.ReleaseMode)
	gin.Recovery()

	// ヘルスチェックエンドポイント
	g.router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	// ルーターの登録
	v1 := g.router.Group("/v1")
	for _, r := range g.routers {
		r.Register(v1)
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := g.server.ListenAndServe(); err != nil {
			fmt.Println(err)
		}
	}()

	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()

	if err := g.server.Shutdown(ctx); err != nil {
		fmt.Println(err)
	}
}
