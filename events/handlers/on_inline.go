package handlers

import (
	"log"

	"github.com/amarnathcjd/gogram/telegram"
)

func HandleOnInline(message *telegram.InlineQuery) error {
	log.Printf("Получен альбом с ID сообщений: %v", message.Sender)
	return nil
}
