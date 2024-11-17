package main

import (
	"log"
	"post-wizard/config"

	"github.com/amarnathcjd/gogram/telegram"
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
			if message.Channel.ID == channelID {
				// Определяем тип сообщения
				if message.Message.Media != nil {
					// Пересылка медиа-сообщения
					forwardMedia(client, message, cfg.TargetUserIDs)
				} else {
					// Пересылка текстового сообщения
					forwardText(client, message, cfg.TargetUserIDs)
				}
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
			Silent: false,
		})
		if err != nil {
			log.Printf("Ошибка при пересылке медиа-сообщения пользователю %d: %v", userID, err)
		} else {
			log.Printf("Медиа-сообщение успешно переслано пользователю %d", userID)
		}
	}
}
