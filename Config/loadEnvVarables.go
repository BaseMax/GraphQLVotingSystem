package Config

import (
	"log"
	"github.com/joho/godotenv"
)

func LoadEnvVarables() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
