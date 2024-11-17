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

	appIDStr := os.Getenv("APP_ID")
	appHash := os.Getenv("APP_HASH")
	phoneNumber := os.Getenv("PHONE_NUMBER")
	channelIDsStr := os.Getenv("CHANNEL_IDS")
	targetUserIDsStr := os.Getenv("TARGET_USER_IDS")

	if appIDStr == "" || appHash == "" || phoneNumber == "" || channelIDsStr == "" || targetUserIDsStr == "" {
		log.Fatal("Не удалось получить конфигурационные параметры из .env")
	}

	appID, err := strconv.Atoi(appIDStr)
	if err != nil {
		log.Fatalf("Ошибка преобразования APP_ID в число: %v", err)
	}

	// Преобразуем списки ID каналов и пользователей в массивы
	channelIDs := parseIDs(channelIDsStr)
	targetUserIDs := parseIDs(targetUserIDsStr)

	return &Config{
		AppID:         int32(appID),
		AppHash:       appHash,
		PhoneNumber:   phoneNumber,
		ChannelIDs:    channelIDs,
		TargetUserIDs: targetUserIDs,
	}
}

// parseIDs преобразует строку "1,2,3" в []int64
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
