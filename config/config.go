package config

import (
	"github.com/go-kit/kit/log"
	"github.com/jinzhu/configor"
)

var Config = struct {
	Database struct {
		Postgres struct {
			Host     string `default:"localhost" env:"DATABASE__POSTGRES__HOST"`
			Port     string `default:"5432" env:"DATABASE__POSTGRES__PORT"`
			User     string `default:"postgres" env:"DATABASE__POSTGRES__USER"`
			Password string `default:"changeme" env:"DATABASE__POSTGRES__PASSWORD"`
			Dbname   string `default:"todo_db" env:"DATABASE__POSTGRES__DBNAME"`
		}
	}
}{}

func LoadConfig(logger log.Logger) error {
	logger.Log("msg", "Load Application Configuration")
	return configor.Load(&Config)
}
