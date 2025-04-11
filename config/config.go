package config

import (
	"github.com/dailoi280702/vrs-general-service/log"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

var (
	defaultPort = "9000"
	config      *Config
)

type StageType string

type Config struct {
	Port    string `envconfig:"SERVER_PORT"`
	IsDebug bool   `envconfig:"DEBUG"`

	MySQL struct {
		Host   string `envconfig:"DB_HOST"`
		Port   string `envconfig:"DB_PORT"`
		User   string `envconfig:"DB_USER"`
		Pass   string `envconfig:"DB_PASS"`
		DBName string `envconfig:"DB_NAME"`
	}

	Redis struct {
		Host string `envconfig:"REDIS_HOST"`
		Port string `envconfig:"REDIS_PORT"`
		DB   int    `envconfig:"REDIS_DB"`
		User string `envconfig:"REDIS_USER"`
		Pass string `envconfig:"REDIS_PASS"`
	}
}

func init() {
	config = &Config{}

	if err := godotenv.Load(); err != nil {
		log.Logger().Warn(err.Error())
	}

	if err := envconfig.Process("", config); err != nil {
		log.Logger().Error(err.Error())
	}

	if len(config.Port) == 0 {
		config.Port = defaultPort
	}
}

func GetConfig() *Config {
	return config
}
