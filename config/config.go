package config

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	AppID         int32
	AppHash       string
	PhoneNumber   string
	ChannelIDs    []int64
	TargetUserIDs []int64
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Не удалось загрузить .env файл, используем переменные окружения из системы")
	}

	return &Config{
		AppID:         getEnvAsInt32("APP_ID"),
		AppHash:       getEnv("APP_HASH"),
		PhoneNumber:   getEnv("PHONE_NUMBER"),
		ChannelIDs:    parseIDs(getEnv("CHANNEL_IDS")),
		TargetUserIDs: parseIDs(getEnv("TARGET_USER_IDS")),
	}
}

func getEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("Отсутствует обязательная переменная окружения: %s", key)
	}
	return value
}

func getEnvAsInt32(key string) int32 {
	valueStr := getEnv(key)
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		log.Fatalf("Ошибка преобразования %s в число: %v", key, err)
	}
	return int32(value)
}

func parseIDs(idsStr string) []int64 {
	ids := []int64{}
	for _, id := range strings.Split(idsStr, ",") {
		parsedID, err := strconv.ParseInt(strings.TrimSpace(id), 10, 64)
		if err == nil {
			ids = append(ids, parsedID)
		}
	}
	return ids
}
