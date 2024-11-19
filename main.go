package main

import (
	"log"
	"post-wizard/config"
	"post-wizard/events"

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

	// Инициализация событий
	events.InitEventHandlers(client, cfg)

	client.Idle()
}
