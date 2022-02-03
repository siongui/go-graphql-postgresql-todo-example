package config

import (
	"github.com/jinzhu/configor"
)

var Config = struct {
	App struct {
		Port string `default:"3005" env:"APP__PORT"`
	}

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

func LoadConfig() error {
	return configor.Load(&Config)
}
