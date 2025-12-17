package config

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	// Try loading from current directory and parent directory
	// This works whether running from backend/ or backend/cmd/
	err := godotenv.Load(".env")
	if err != nil {
		// Try parent directory (for when running from cmd/)
		err = godotenv.Load("../.env")
		if err != nil {
			log.Printf("Warning: Could not load .env file. Make sure .env exists in the backend directory.")
			log.Println("Continuing without .env file. Using system environment variables.")
		}
	}
}
