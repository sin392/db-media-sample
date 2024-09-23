package database

import (
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
		Host:       "localhost",
		Port:       "27017",
		Database:   "test",
		Driver:     "mongodb",
		User:       "root",
		Password:   "password",
		CtxTimeout: 60 * time.Second,
	}
}
