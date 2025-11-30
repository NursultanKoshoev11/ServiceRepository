package config

import (
	"os"
)

type Config struct {
	DBDriver      string
	DBHost        string
	DBPort        string
	DBUser        string
	DBPassword    string
	DBName        string
	MigrationPath string
	GRPCPort      string
}

func LoadConfig() *Config {
	return &Config{
		DBDriver:      getEnv("DB_DRIVER", "postgres"),
		DBHost:        getEnv("DB_HOST", "localhost"),
		DBPort:        getEnv("DB_PORT", "5432"),
		DBUser:        getEnv("DB_USER", "postgres"),
		DBPassword:    getEnv("DB_PASSWORD", "123"),
		DBName:        getEnv("DB_NAME", "testdb"),
		MigrationPath: getEnv("DB_MIGRATIONPATH", "file://internal/migrations"),
		GRPCPort:      getEnv("GRPCPort", ":50051"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
