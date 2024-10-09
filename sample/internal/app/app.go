package app

import (
	"sync"

	"github.com/sin392/db-media-sample/sample/internal/config"
	"github.com/sin392/db-media-sample/sample/internal/infrastructure/server"
	"github.com/sin392/db-media-sample/sample/module/trace"
)

type Application struct {
	cfg        *config.Config
	grpcServer *server.GrpcServer
	httpServer *server.HttpServer
	gqlServer  *server.GqlServer
	wg         sync.WaitGroup
}

func NewApplication(
	cfg *config.Config,
	grpcServer *server.GrpcServer,
	httpServer *server.HttpServer,
	gqlServer *server.GqlServer,
) (*Application, error) {
	app := &Application{
		cfg:        cfg,
		grpcServer: grpcServer,
		httpServer: httpServer,
		gqlServer:  gqlServer,
		wg:         sync.WaitGroup{},
	}
	app.configure()
	return app, nil
}

func (a *Application) configure() {
	// OpenTelemetryの初期化
	trace.InitTraceProvider(a.cfg)
}

func (a *Application) Start() {
	grpcServerEndpoint := ":50051"
	httpServerEndpoint := ":8080"
	gqlServerEndpoint := ":8081"

	a.wg.Add(3) // 非同期に起動するサーバの数だけカウントアップ
	go func() {
		a.grpcServer.ListenAndServe(grpcServerEndpoint)
		a.wg.Done()
	}()
	go func() {
		conn, err := server.NewGrpcConnection(grpcServerEndpoint)
		if err != nil {
			panic(err)
		}
		a.httpServer.ListenAndServe(httpServerEndpoint, conn)
		a.wg.Done()
	}()
	go func() {
		conn, err := server.NewGrpcConnection(grpcServerEndpoint)
		if err != nil {
			panic(err)
		}
		a.gqlServer.ListenAndServe(gqlServerEndpoint, conn)
		a.wg.Done()
	}()
	a.wg.Wait()
}

// TODO: 停止処理の追加
func (a *Application) Stop() {
	// a.grpcServer.Stop()
	// a.httpServer.Stop()
}
