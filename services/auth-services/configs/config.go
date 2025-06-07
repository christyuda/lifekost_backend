package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	JWTSecret string
	DBUrl     string
	Port      string
}

var AppConfig *Config

func LoadConfig() {
	// Load .env file (optional, ignore error if not found)
	_ = godotenv.Load()

	AppConfig = &Config{
		JWTSecret: getEnv("JWT_SECRET", "default_secret"),
		DBUrl:     getEnv("DB_URL", "postgres://user:pass@localhost:5432/authdb?sslmode=disable"),
		Port:      getEnv("PORT", "8001"),
	}

	log.Println("✅ Configuration loaded")
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
