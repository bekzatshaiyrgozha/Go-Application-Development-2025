// main.go
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// Версия и метаданные приложения - будут заполнены при сборке
var (
	Version   = "dev"
	BuildTime = "unknown"
	GitCommit = "unknown"
)

// Простая структура для хранения данных пользователя
type User struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func main() {
	// Вывод информации о версии
	fmt.Printf("Starting application version %s (build: %s, commit: %s)\n",
		Version, BuildTime, GitCommit)

	// Простой HTTP-сервер
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/api/users", usersHandler)
	http.HandleFunc("/api/health", healthHandler)

	// Получение порта из переменной окружения или использование значения по умолчанию
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Запуск сервера
	addr := ":" + port
	fmt.Printf("Server is running on port %s\n", port)
	log.Fatal(http.ListenAndServe(addr, nil))
}

// Обработчик главной страницы
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Demo API!")
}

// Обработчик для API пользователей
func usersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `[{"id":"1","username":"user1","email":"user1@example.com"}]`)
}

// Обработчик для проверки здоровья сервиса
func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"status":"ok","timestamp":"%s","version":"%s"}`,
		time.Now().Format(time.RFC3339), Version)
}
