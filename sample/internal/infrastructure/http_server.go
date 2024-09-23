package infrastructure

import (
	"strconv"
	"time"

	"github.com/sin392/db-media-sample/internal/adapter/repositoryimpl/nosql"
	"github.com/sin392/db-media-sample/internal/infrastructure/database"
	"github.com/sin392/db-media-sample/internal/infrastructure/router"
)

type config struct {
	appName       string
	dbNoSQL       nosql.NoSQL
	ctxTimeout    time.Duration
	webServerPort router.Port
	webServer     router.Server
}

func NewConfig() *config {
	return &config{}
}

func (c *config) ContextTimeout(t time.Duration) *config {
	c.ctxTimeout = t
	return c
}

func (c *config) Name(name string) *config {
	c.appName = name
	return c
}

func (c *config) DbNoSQL() *config {
	db, err := database.NewMongoHandler(database.NewConfig())
	if err != nil {
		return nil
	}

	c.dbNoSQL = db
	return c
}

func (c *config) WebServer() *config {
	s := router.NewGinServer(
		c.dbNoSQL,
		c.webServerPort,
		c.ctxTimeout,
	)

	c.webServer = s
	return c
}

func (c *config) WebServerPort(port string) *config {
	p, err := strconv.ParseInt(port, 10, 64)
	if err != nil {
		return nil
	}

	c.webServerPort = router.Port(p)
	return c
}

func (c *config) Start() {
	c.webServer.Listen()
}
