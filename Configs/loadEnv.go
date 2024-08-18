package Configs

import (
	"github.com/joho/godotenv"
	"log"
)

func LoadENV() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}
