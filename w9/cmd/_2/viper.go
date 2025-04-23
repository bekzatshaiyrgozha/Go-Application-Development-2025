package _2

import (
	"github.com/spf13/viper"
)

type ViperConfig struct {
	Server struct {
		Port    string
		Timeout int
	}
	Database struct {
		Host     string
		Port     string
		Username string
		Password string
		Name     string
	}
	Auth struct {
		JwtSecret string
		TokenTTL  int
	}
}

func LoadViperConfig() (ViperConfig, error) {
	var config ViperConfig

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	viper.AddConfigPath(".")

	// Поддержка переменных окружения с префиксом APP
	viper.SetEnvPrefix("APP")
	viper.AutomaticEnv()

	// Сопоставление переменных окружения с конфигурацией
	viper.BindEnv("server.port", "APP_SERVER_PORT")
	viper.BindEnv("database.host", "APP_DB_HOST")
	// ... другие переменные

	// Установка значений по умолчанию
	viper.SetDefault("server.port", "8080")
	viper.SetDefault("server.timeout", 30)

	if err := viper.ReadInConfig(); err != nil {
		return config, err
	}

	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}

	return config, nil
}
