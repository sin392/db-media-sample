package config

import "time"

type Config struct {
	AppName       string        `mapstructure:"app_name"`
	AppVersion    string        `mapstructure:"app_version"`
	WebServerPort int           `mapstructure:"web_server_port"`
	Timeout       time.Duration `mapstructure:"timeout"`
}
