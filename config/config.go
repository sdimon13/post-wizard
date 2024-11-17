package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	AppID       int32
	AppHash     string
	PhoneNumber string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Не удалось загрузить .env файл, используем переменные окружения из системы")
	}

	appIDStr := os.Getenv("APP_ID")
	appHash := os.Getenv("APP_HASH")
	phoneNumber := os.Getenv("PHONE_NUMBER")

	if appIDStr == "" || appHash == "" || phoneNumber == "" {
		log.Fatal("Не удалось получить APP_ID, APP_HASH или PHONE_NUMBER из переменных окружения")
	}

	// Преобразуем APP_ID из строки в int32
	appID, err := strconv.Atoi(appIDStr)
	if err != nil {
		log.Fatalf("Ошибка преобразования APP_ID в число: %v", err)
	}

	return &Config{
		AppID:       int32(appID),
		AppHash:     appHash,
		PhoneNumber: phoneNumber,
	}
}
