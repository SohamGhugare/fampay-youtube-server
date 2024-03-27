package initializers

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	// loading the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load .env file...")
	}

	// checking if API key exists in .env file
	_, keyExists := os.LookupEnv("API_KEY")
	if !keyExists {
		log.Fatal("API_KEY is not set in the .env file...")
	}

}
