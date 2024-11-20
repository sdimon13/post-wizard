package main

import (
	"fmt"
	"github.com/amarnathcjd/gogram/telegram"
	"log"
	"post-wizard/config"
)

func main() {
	cfg := config.LoadConfig()

	client, err := telegram.NewClient(telegram.ClientConfig{
		AppID:    cfg.AppID,
		AppHash:  cfg.AppHash,
		LogLevel: telegram.LogInfo,
	})

	if err != nil {
		log.Fatal(err)
	}

	client.Conn()

	client.Login(cfg.PhoneNumber)

	log.Println(cfg.ChannelIDs)

	// Подписываемся на события сообщений
	client.On(telegram.OnMessage, func(message *telegram.NewMessage) error {
		log.Printf("Received message from id: %d, name: %s\n", message.Sender.ID, message.Sender.Username)

		// Игнорируем сообщения, не пришедшие из каналов
		if message.Channel == nil {
			return nil
		}

		for _, channelID := range cfg.ChannelIDs {
			if message.Channel.ID == channelID && message.Message.GroupedID == 0 {
				forwardMedia(client, message, cfg.TargetUserIDs)
				break
			}
		}
		return nil
	})

	client.Idle()
}

func forwardText(client *telegram.Client, message *telegram.NewMessage, targetUserIDs []int64) {
	for _, userID := range targetUserIDs {
		_, err := client.SendMessage(userID, message.Message.Message)
		if err != nil {
			log.Printf("Ошибка при отправке текстового сообщения пользователю %d: %v", userID, err)
		} else {
			log.Printf("Текстовое сообщение успешно переслано пользователю %d", userID)
		}
	}
}

func forwardMedia(client *telegram.Client, message *telegram.NewMessage, targetUserIDs []int64) {
	for _, userID := range targetUserIDs {
		_, err := client.Forward(userID, message.Channel.ID, []int32{message.ID}, &telegram.ForwardOptions{
			HideAuthor: true,
		})
		if err != nil {
			log.Printf("Ошибка при пересылке медиа-сообщения пользователю %d: %v", userID, err)
		} else {
			log.Printf("Медиа-сообщение успешно переслано пользователю %d", userID)
		}
	}
}

func forwardMessage(client *telegram.Client, message *telegram.NewMessage, targetUserIDs []int64) {
	for _, userID := range targetUserIDs {
		_, err := client.SendMessage(userID, &telegram.NewMessage{
			Message: message.Message,
		})
		if err != nil {
			log.Printf("Ошибка при отправке текстового сообщения пользователю %d: %v", userID, err)
		} else {
			log.Printf("Текстовое сообщение успешно переслано пользователю %d", userID)
		}
	}
}

func saveMedia(client *telegram.Client, message *telegram.NewMessage, targetUserIDs []int64) {
	options := &telegram.DownloadOptions{
		Threads: 4,
		ProgressCallback: func(totalBytes int64, downloadedBytes int64) {
			fmt.Printf("Загружено %d из %d байт\n", downloadedBytes, totalBytes)
		},
	}
	// Загрузка медиа-контента напрямую на диск
	savedFilePath, err := message.Download(options)
	if err != nil {
		fmt.Println("Ошибка загрузки медиа:", err)
		return
	}
	fmt.Println("Медиа успешно загружено в:", savedFilePath)
}
