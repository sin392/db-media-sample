package app

import (
	"sync"

	"github.com/sin392/db-media-sample/sample/internal/config"
	"github.com/sin392/db-media-sample/sample/internal/infrastructure"
	"github.com/sin392/db-media-sample/sample/module/trace"
)

type Application struct {
	cfg        *config.Config
	grpcServer infrastructure.GrpcServer
	httpServer infrastructure.HttpServer
	wg         sync.WaitGroup
}

func NewApplication(
	cfg *config.Config,
	grpcServer infrastructure.GrpcServer,
	httpServer infrastructure.HttpServer,
) (*Application, error) {
	return &Application{
		cfg:        cfg,
		grpcServer: grpcServer,
		httpServer: httpServer,
		wg:         sync.WaitGroup{},
	}, nil
}

func (a *Application) Configure() {
	// OpenTelemetryの初期化
	trace.InitTraceProvider(a.cfg)
}

func (a *Application) Start() {
	a.wg.Add(2) // 非同期に起動するサーバの数だけカウントアップ
	go func() {
		a.grpcServer.ListenAndServe()
		a.wg.Done()
	}()
	go func() {
		a.httpServer.ListenAndServe()
		a.wg.Done()
	}()
	a.wg.Wait()
}

// TODO: 停止処理の追加
func (a *Application) Stop() {
	// a.grpcServer.Stop()
	// a.httpServer.Stop()
}
