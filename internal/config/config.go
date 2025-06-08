package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Env           string
	Port          string
	DBUrl         string
	AllowedOrigin string
}

func Load() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, using environment variables")
	}

	return &Config{
		Env:           getEnv("ENV", "production"),
		Port:          getEnv("PORT", "8080"),
		DBUrl:         getEnv("DB_URL", "storage/puppet.db"),
		AllowedOrigin: getEnv("ALLOWED_ORIGIN", "http://localhost:5173"),
	}
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
