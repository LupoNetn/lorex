package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port  string // Capitalized to make them exported
	DBUrl string
}

func LoadConfig() (*Config, error) {
	// It's often okay if .env is missing in production (env vars might be set on the host)
	// but for local dev, this check is fine.
	if err := godotenv.Load(); err != nil {
		return nil, errors.New("no .env file found")
	}

	// Logic to populate the struct
	port, err := ExtractEnvKey("PORT", "8080")
	if err != nil {
		return nil, err
	}

	dbUrl, err := ExtractEnvKey("DATABASE_URL", "")
	if err != nil {
		return nil, err
	}

	return &Config{
		Port:  port,
		DBUrl: dbUrl,
	}, nil
}

func ExtractEnvKey(key, fallback string) (string, error) {
	value := os.Getenv(key)
	if value == "" {
		if fallback != "" {
			return fallback, nil
		}
		return "", errors.New("required env var missing: " + key)
	}
	return value, nil
}