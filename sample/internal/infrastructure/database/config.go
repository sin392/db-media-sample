package database

import (
	"os"
	"time"
)

type DBConfig struct {
	Host     string
	Database string
	Port     string
	Driver   string
	User     string
	Password string

	CtxTimeout time.Duration
}

func NewConfig() *DBConfig {
	return &DBConfig{
		Host:       os.Getenv("MONGODB_HOST"),
		Port:       os.Getenv("MONGODB_PORT"),
		Database:   "test",
		Driver:     "mongodb",
		User:       "root",
		Password:   "password",
		CtxTimeout: 60 * time.Second,
	}
}
