package handlers

import (
	"log"
	"post-wizard/config"
	"post-wizard/services"

	"github.com/amarnathcjd/gogram/telegram"
)

func HandleOnMessage(client *telegram.Client, message *telegram.NewMessage, cfg *config.Config) error {
	log.Printf("Получено сообщение от: %s (ID: %d)", message.Sender.Username, message.Sender.ID)

	if message.Message.Message != "" {
		services.ForwardText(client, message, cfg.TargetUserIDs)
	} else {
		services.ForwardMedia(client, message, cfg.TargetUserIDs)
	}

	return nil
}
