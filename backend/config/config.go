package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("No .env file found, relying on system environment variables")
	}
}

func GetEnv(key string) string {
	return os.Getenv(key)
}
