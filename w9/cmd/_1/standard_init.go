package _1

import (
	"log"
	"net/http"
)

func main() {
	// Инициализация компонентов
	config := loadConfig()
	db := initDatabase(config)

	// Создание и настройка роутера
	router := setupRouter(db)

	// Запуск сервера
	log.Printf("Starting server on :%s", config.Port)
	if err := http.ListenAndServe(":"+config.Port, router); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
