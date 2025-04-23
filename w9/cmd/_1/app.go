package _1

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
)

type App struct {
	config Config
	db     *sql.DB
	server *http.Server
	logger *log.Logger
	// Другие зависимости
}

// Конструктор приложения
func NewApp(config Config) *App {
	return &App{
		config: config,
		logger: log.New(os.Stdout, "[APP] ", log.LstdFlags),
	}
}

func (a *App) Start(ctx context.Context) error {
	// Инициализация компонентов
	if err := a.initDatabase(); err != nil {
		return err
	}

	a.initRouter()

	// Запуск HTTP сервера
	go a.startServer()

	// Ожидание сигнала завершения
	<-ctx.Done()

	// Graceful shutdown
	return a.shutdown()
}
