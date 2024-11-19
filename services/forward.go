package services

import (
	"log"

	"github.com/amarnathcjd/gogram/telegram"
)

func ForwardText(client *telegram.Client, message *telegram.NewMessage, targetUserIDs []int64) {
	for _, userID := range targetUserIDs {
		_, err := client.SendMessage(userID, message.Message.Message)
		if err != nil {
			log.Printf("Ошибка при отправке текста пользователю %d: %v", userID, err)
		} else {
			log.Printf("Текст успешно отправлен пользователю %d", userID)
		}
	}
}

func ForwardMedia(client *telegram.Client, message *telegram.NewMessage, targetUserIDs []int64) {
	for _, userID := range targetUserIDs {
		_, err := client.Forward(userID, message.Channel.ID, []int32{message.ID}, &telegram.ForwardOptions{
			HideAuthor: true,
		})
		if err != nil {
			log.Printf("Ошибка при пересылке медиа пользователю %d: %v", userID, err)
		} else {
			log.Printf("Медиа успешно переслано пользователю %d", userID)
		}
	}
}

func SendAlbum(client *telegram.Client, album *telegram.Album, targetUserIDs []int64) {
	for _, userID := range targetUserIDs {
		_, err := client.SendAlbum(userID, album)
		if err != nil {
			log.Printf("Ошибка при отправке текста пользователю %d: %v", userID, err)
		} else {
			log.Printf("Текст успешно отправлен пользователю %d", userID)
		}
	}
}
