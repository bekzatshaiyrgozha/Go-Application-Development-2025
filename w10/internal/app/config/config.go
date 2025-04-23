package config

import (
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	"github.com/mcuadros/go-defaults"
)

type Config struct {
	HTTPServer HTTPServerConfig
	DB         DBConfig
}

type HTTPServerConfig struct {
	Addr           string `env:"HTTP_ADDR" default:":8080"`
	HealthAddr     string `env:"HTTP_HEALTH_ADDR" default:":8081"`
	RequestTimeout int    `env:"HTTP_REQUEST_TIMEOUT" default:"60"`
}

type DBConfig struct {
	Host     string `env:"POSTGRES_HOST" default:"localhost"`
	Port     string `env:"POSTGRES_PORT" default:"5432"`
	User     string `env:"POSTGRES_USER" default:"dbuser"`
	Password string `env:"POSTGRES_PASSWORD" default:"dbpassword"`
	Database string `env:"POSTGRES_DB" default:"servicedb"`
	SSLMode  string `env:"POSTGRES_SSL_MODE" default:"disable"`
}

func New(filenames ...string) (*Config, error) {
	cfg := new(Config)

	if len(filenames) > 0 {
		if err := godotenv.Load(filenames...); err != nil {
			return nil, err
		}
	}

	if err := env.Parse(cfg); err != nil {
		return nil, err
	}

	defaults.SetDefaults(cfg)

	return cfg, nil
}
