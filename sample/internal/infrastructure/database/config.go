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
		database:   os.Getenv("MONGODB_DATABASE"),
		driver:     os.Getenv("MONGODB_DRIVER"),
		user:       os.Getenv("MONGODB_ROOT_USER"),
		password:   os.Getenv("MONGODB_ROOT_PASSWORD"),
		ctxTimeout: 60 * time.Second,
	}
}
