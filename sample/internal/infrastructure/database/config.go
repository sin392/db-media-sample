package database

import (
	"os"
	"time"
)

type config struct {
	host     string
	database string
	port     string
	driver   string
	user     string
	password string

	ctxTimeout time.Duration
}

func NewConfig() *config {
	return &config{
		host:       os.Getenv("MONGODB_HOST"),
		port:       os.Getenv("MONGODB_PORT"),
		database:   "test",
		driver:     "mongodb",
		user:       "root",
		password:   "password",
		ctxTimeout: 60 * time.Second,
	}
}
