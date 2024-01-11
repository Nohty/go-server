package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	PORT       string
	JWT_SECRET string
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	if fallback == "" {
		log.Panicln("Environment variable required but not set: ", key)
	}

	return fallback
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	PORT = getEnv("PORT", "3000")
	JWT_SECRET = getEnv("JWT_SECRET", "")
}
