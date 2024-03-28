package database

import (
	"log"

	"github.com/SohamGhugare/fampay-youtube-server/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Client *gorm.DB

func SetupDatabaseClient() {
	var err error
	Client, err = gorm.Open(sqlite.Open("data/videos.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("An error occured while setting up database client: %v", err)
	}

	Client.AutoMigrate(&models.Video{})

	log.Println("Database client setup successful...")
}
