package events

import (
	"post-wizard/config"
	"post-wizard/events/handlers"

	"github.com/amarnathcjd/gogram/telegram"
)

func InitEventHandlers(client *telegram.Client, cfg *config.Config) {
	client.On(telegram.OnMessage, func(message *telegram.NewMessage) error {
		return handlers.HandleOnMessage(client, message, cfg)
	}, telegram.FilterFunc(func(message *telegram.NewMessage) bool {
		return message.Message.Post
	}))
	/*client.On(telegram.OnAlbum, func(album *telegram.Album) error {
		return handlers.HandleOnAlbum(client, album, cfg)
	}, telegram.FilterFunc(func(message *telegram.NewMessage) bool {
		return message.Message.InvertMedia
	}))*/
	client.On(telegram.OnInline, handlers.HandleOnInline)
}
