package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	// Server configuration
	ServerHost string `env:"SERVER_HOST"`
	ServerPort string `env:"SERVER_PORT"`

	// TLS configuration
	TLSCertFilePath string `env:"TLS_CERT_FILE_PATH"`
	TLSKeyFilePath  string `env:"TLS_KEY_FILE_PATH"`

	// Database configuration
	DBHost     string `env:"DB_HOST"`
	DBPort     string `env:"DB_PORT"`
	DBUser     string `env:"DB_USER"`
	DBPassword string `env:"DB_PASSWORD"`
	DBName     string `env:"DB_NAME"`

	// Environment
	Environment string `env:"ENVIRONMENT"`

	// Logging
	LogLevel string `env:"LOG_LEVEL"`
}

func NewConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		if !os.IsNotExist(err) {
			return nil, fmt.Errorf("error loading .env file: %w", err)
		}
	}

	cfg := &Config{}

	cfg.ServerHost = getEnv("SERVER_HOST", "0.0.0.0")
	cfg.ServerPort = getEnv("SERVER_PORT", "8080")
	cfg.TLSCertFilePath = getEnv("TLS_CERT_FILE_PATH", "")
	cfg.TLSKeyFilePath = getEnv("TLS_KEY_FILE_PATH", "")
	cfg.DBHost = getEnv("DB_HOST", "localhost")
	cfg.DBPort = getEnv("DB_PORT", "5432")
	cfg.DBUser = getEnv("DB_USER", "postgres")
	cfg.DBPassword = getEnv("DB_PASSWORD", "")
	cfg.DBName = getEnv("DB_NAME", "auth_db")
	cfg.Environment = getEnv("ENVIRONMENT", "development")
	cfg.LogLevel = getEnv("LOG_LEVEL", "info")

	return cfg, nil
}

func getEnv(key string, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
