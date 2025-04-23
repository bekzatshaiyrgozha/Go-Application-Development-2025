package _2

import (
	"flag"
)

type FlagConfig struct {
	Port        string
	DatabaseURL string
	LogLevel    string
	Environment string
}

func LoadFlagConfig() FlagConfig {
	config := FlagConfig{}

	flag.StringVar(&config.Port, "port", "8080", "Server port")
	flag.StringVar(&config.DatabaseURL, "db-url", "postgres://user:pass@localhost:5432/appdb", "Database connection string")
	flag.StringVar(&config.LogLevel, "log-level", "info", "Logging level")
	flag.StringVar(&config.Environment, "env", "development", "Environment (development/staging/production)")

	flag.Parse()

	return config
}
