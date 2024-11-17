package main

import (
	"github.com/amarnathcjd/gogram/telegram"
	"log"
	"post-wizard/config"
)

func main() {
	cfg := config.LoadConfig()

	client, err := telegram.NewClient(telegram.ClientConfig{
		AppID:   cfg.AppID,
		AppHash: cfg.AppHash,
	})

	if err != nil {
		log.Fatal(err)
	}

	client.Conn()

	client.Login(cfg.PhoneNumber)

	client.On(telegram.OnMessage, func(message *telegram.NewMessage) error {
		message.Reply("Hello from Gogram!")
		return nil
	}, telegram.FilterPrivate)

	client.Idle()
}
