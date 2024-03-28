package database

import (
	"log"
	"time"

	"github.com/SohamGhugare/fampay-youtube-server/services"
)

func Loop() {
	for {
		log.Println("----------------------------------")
		log.Println("Looping...")

		// clearing the database before fetching new videos
		Client.Exec("DELETE FROM videos")

		// fetching and saving 10 latest videos with query golang
		videos := services.FetchSvc("golang", 10)
		for _, video := range videos {
			Client.Create(&video)
			log.Printf("Saved Video [%v]: %v\n", video.ID, video.Title)
		}
		log.Println("---------------------------------")
		time.Sleep(10 * time.Second)

	}
}
