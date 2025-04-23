package _1

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Создаем контекст, который будет отменен при получении сигнала
	ctx, cancel := context.WithCancel(context.Background())

	// Настраиваем обработку сигналов для graceful shutdown
	go func() {
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
		<-sigCh
		log.Println("Shutdown signal received")
		cancel()
	}()

	// Инициализируем и запускаем приложение
	app := NewApp(LoadConfig())
	if err := app.Start(ctx); err != nil {
		log.Fatalf("Failed to start application: %v", err)
	}
}
