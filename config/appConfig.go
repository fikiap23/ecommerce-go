package config

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	ServerPort string
	Dsn        string
	AppSecret  string
}

// SetupEnv loads environment variables, optionally from a .env file
func SetupEnv(dotenvPath string) (AppConfig, error) {
	// Determine if we should load from a .env file
	appEnv := strings.TrimSpace(os.Getenv("APP_ENV"))
	shouldLoadDotenv := appEnv == "dev" || dotenvPath != ""

	// Load .env file if needed
	if shouldLoadDotenv {
		envFile := ".env"
		if dotenvPath != "" {
			envFile = dotenvPath
		}

		if err := godotenv.Load(envFile); err != nil {
			return AppConfig{}, fmt.Errorf("failed to load env file '%s': %w", envFile, err)
		}
	}

	// Fetch required environment variables
	httpPort := strings.TrimSpace(os.Getenv("HTTP_PORT"))
	if httpPort == "" {
		return AppConfig{}, errors.New("missing required environment variable: HTTP_PORT")
	}

	dsn := strings.TrimSpace(os.Getenv("DSN"))
	if dsn == "" {
		return AppConfig{}, errors.New("missing required environment variable: DSN")
	}

	appSecret := strings.TrimSpace(os.Getenv("APP_SECRET"))
	if appSecret == "" {
		return AppConfig{}, errors.New("missing required environment variable: APP_SECRET")
	}

	return AppConfig{
		ServerPort: httpPort,
		Dsn:        dsn,
		AppSecret:  appSecret,
	}, nil
}

