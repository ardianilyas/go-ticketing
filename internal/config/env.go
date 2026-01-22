package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("failed to load environment variables", err)
	}
}

func Get(key string) string {
	return os.Getenv(key)
}