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
		log.Printf("Received message from id: %d, name: %s: %s\n", message.Sender.ID, message.Sender.Username, message.Message.Message)

		// Игнорируем сообщения, не пришедшие из каналов
		if message.Channel == nil {
			return nil
		}

		for _, channelID := range cfg.ChannelIDs {
			if message.Channel.ID == channelID {
				// Пересылаем сообщение всем указанным пользователям
				for _, userID := range cfg.TargetUserIDs {
					_, err := client.SendMessage(userID, message.Message.Message)

					if err != nil {
						log.Printf("Ошибка при отправке сообщения пользователю %d: %v", userID, err)
					} else {
						log.Printf("Сообщение переслано пользователю %d", userID)
					}
				}
				break // Не нужно повторять цикл для других каналов
			}
		}

		return nil
	})

	client.Idle()
}
