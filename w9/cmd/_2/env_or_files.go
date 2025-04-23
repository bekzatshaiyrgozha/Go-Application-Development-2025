package _2

import (
	"encoding/json"
	"io"
	"os"
)

type Config struct {
	Server struct {
		Port    string `json:"port"`
		Timeout int    `json:"timeout"`
	} `json:"server"`
	Database struct {
		Host     string `json:"host"`
		Port     string `json:"port"`
		Username string `json:"username"`
		Password string `json:"password"`
		Name     string `json:"name"`
	} `json:"database"`
	Auth struct {
		JwtSecret string `json:"jwt_secret"`
		TokenTTL  int    `json:"token_ttl"`
	} `json:"auth"`
}

func LoadConfig(filePath string) (Config, error) {
	var config Config

	// Загружаем настройки из файла
	file, err := os.Open(filePath)
	if err != nil {
		return config, err
	}
	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		return config, err
	}

	if err := json.Unmarshal(byteValue, &config); err != nil {
		return config, err
	}

	// Переопределяем значениями из окружения, если они есть
	if port := os.Getenv("SERVER_PORT"); port != "" {
		config.Server.Port = port
	}

	if dbHost := os.Getenv("DB_HOST"); dbHost != "" {
		config.Database.Host = dbHost
	}

	return config, nil
}
