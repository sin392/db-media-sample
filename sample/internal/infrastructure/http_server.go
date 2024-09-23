package infrastructure

import (
	"strconv"
	"time"

	"github.com/sin392/db-media-sample/internal/adapter/repositoryimpl/nosql"
	"github.com/sin392/db-media-sample/internal/infrastructure/database"
	"github.com/sin392/db-media-sample/internal/infrastructure/router"
)

type Config struct {
	appName       string
	dbNoSQL       nosql.NoSQL
	ctxTimeout    time.Duration
	webServerPort router.Port
	webServer     router.Server
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) ContextTimeout(t time.Duration) *Config {
	c.ctxTimeout = t
	return c
}

func (c *Config) Name(name string) *Config {
	c.appName = name
	return c
}

func (c *Config) DbNoSQL() *Config {
	db, err := database.NewMongoHandler(database.NewConfig())
	if err != nil {
		return nil
	}

	c.dbNoSQL = db
	return c
}

func (c *Config) WebServer() *Config {
	s := router.NewGinServer(
		c.dbNoSQL,
		c.webServerPort,
		c.ctxTimeout,
	)

	c.webServer = s
	return c
}

func (c *Config) WebServerPort(port string) *Config {
	p, err := strconv.ParseInt(port, 10, 64)
	if err != nil {
		return nil
	}

	c.webServerPort = router.Port(p)
	return c
}

func (c *Config) Start() {
	c.webServer.Listen()
}
