package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system environment variables")
	}
}

var (
	Port       = getEnv("PORT", "5050")
	DBHost     = getEnv("DB_HOST", "db")
	DBPort     = getEnv("DB_PORT", "5432")
	DBUser     = getEnv("DB_USER", "default_user")
	DBPassword = getEnv("DB_PASSWORD", "default_password")
	DBName     = getEnv("DB_NAME", "poetry_db")
	JWTSecret  = getEnv("JWT_SECRET", "supersecret")
)

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
