package internal

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvValues() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
