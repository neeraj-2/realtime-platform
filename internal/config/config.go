package config


import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort        string
	PostgresHost   string
	PostgresPort   string
	PostgresUser   string
	PostgresPassword string
	PostgresDB     string
	RedisAddr      string
}


func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file: %v", err)
		return nil, err
	}

	config := &Config{
		AppPort:        getEnv("APP_PORT", "8080"),
		PostgresHost:   getEnv("POSTGRES_HOST", "localhost"),
		PostgresPort:   getEnv("POSTGRES_PORT", "5432"),
		PostgresUser:   getEnv("POSTGRES_USER", "user"),
		PostgresPassword: getEnv("POSTGRES_PASSWORD", "password"),
		PostgresDB:     getEnv("POSTGRES_DB", "dbname"),
		RedisAddr:      getEnv("REDIS_ADDR", "localhost:6379"),
	}

	return config, nil
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}