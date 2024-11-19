package handlers

import (
	"github.com/amarnathcjd/gogram/telegram"
	"post-wizard/config"
	"post-wizard/services"
)

func HandleOnAlbum(client *telegram.Client, album *telegram.Album, cfg *config.Config) error {
	services.SendAlbum(client, album, cfg.TargetUserIDs)
	return nil
}
