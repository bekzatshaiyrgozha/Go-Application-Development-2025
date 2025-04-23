package _1

import (
	"database/sql"
	"net/http"
)

type Config struct {
	Port string
}

func LoadConfig() Config {
	return Config{}
}

func loadConfig() Config {
	return Config{}
}

func (a *App) initDatabase() error {
	return nil
}

func (a *App) initRouter() {}

func (a *App) startServer() {}

func (a *App) shutdown() error {
	return nil
}

func initDatabase(config Config) *sql.DB {
	return nil
}

func setupRouter(db *sql.DB) http.Handler {
	return nil
}
