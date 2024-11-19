package services

import (
	"fmt"

	"github.com/amarnathcjd/gogram/telegram"
)

func SaveMedia(message *telegram.NewMessage) {
	options := &telegram.DownloadOptions{
		Threads: 4,
		ProgressCallback: func(totalBytes int64, downloadedBytes int64) {
			fmt.Printf("Загружено %d из %d байт\n", downloadedBytes, totalBytes)
		},
	}
	savedFilePath, err := message.Download(options)
	if err != nil {
		fmt.Printf("Ошибка загрузки медиа: %v\n", err)
		return
	}
	fmt.Printf("Медиа успешно загружено в: %s\n", savedFilePath)
}
