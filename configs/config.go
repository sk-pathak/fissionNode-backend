package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Port               string
	DBDriver           string
	DBUser             string
	DBPassword         string
	DBHost             string
	DBPort             string
	DBName             string
	LogLevel           string
	GooseDriver        string
	GooseDBString      string
	GooseMigrationDir  string
	GoogleClientID     string
	GoogleClientSecret string
	GithubClientID     string
	GithubClientSecret string
	SessionSecret      string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, proceeding with environment variables")
	}

	return &Config{
		Port:       getEnv("PORT", "8080"),
		DBDriver:   getEnv("DB_DRIVER", "postgres"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "password"),
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBName:     getEnv("DB_NAME", "postgres"),
		LogLevel:   getEnv("LOG_LEVEL", "info"),

		GooseDriver:       getEnv("GOOSE_DRIVER", "postgres"),
		GooseDBString:     getEnv("GOOSE_DBSTRING", ""),
		GooseMigrationDir: getEnv("GOOSE_MIGRATION_DIR", ""),

		GoogleClientID:     getEnv("GOOGLE_CLIENT_ID", ""),
		GoogleClientSecret: getEnv("GOOGLE_CLIENT_SECRET", ""),

		GithubClientID:     getEnv("GITHUB_CLIENT_ID", ""),
		GithubClientSecret: getEnv("GITHUB_CLIENT_SECRET", ""),

		SessionSecret: getEnv("SESSION_SECRET", "randomString"),
	}, nil
}

func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}
