package config

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv(paths ...string) {
	if len(paths) == 0 {
		paths = []string{".env"}
	}

	err := godotenv.Load(paths...)
	if err != nil {
		log.Printf("No .env file found in paths: %v\n", paths)
	}
}
