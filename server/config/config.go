package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort             string
	DatabaseURL         string
	GoogleClientID      string
	GoogleClientSecret  string
	GoogleRedirectURL   string
	JWTSecret           string
	EncryptionKey       string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, relying on environment variables.")
	}

	return &Config{
		AppPort:             getEnv("APP_PORT", "8080"),
		DatabaseURL:         getEnv("DATABASE_URL", "postgres://user:password@localhost:5432/bicycleapp_db?sslmode=disable"),
		GoogleClientID:      getEnv("GOOGLE_CLIENT_ID", ""),
		GoogleClientSecret:  getEnv("GOOGLE_CLIENT_SECRET", ""),
		GoogleRedirectURL:   getEnv("GOOGLE_REDIRECT_URL", "http://localhost:8080/auth/google/callback"),
		JWTSecret:           getEnv("JWT_SECRET", "supersecretjwtkey"),
		EncryptionKey:       getEnv("ENCRYPTION_KEY", "12345678901234567890123456789012"),
	}
}

func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
