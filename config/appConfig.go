package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	ServerPort       string
	Dsn              string
	AppSecret        string
	TwilioAccountSID string
	TwilioAuthToken  string
	TwilioFromPhone  string
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

	// Fetch and validate required environment variables
	httpPort := getEnv("HTTP_PORT")
	dsn := getEnv("DSN")
	appSecret := getEnv("APP_SECRET")
	twilioSID := getEnv("TWILIO_ACCOUNT_SID")
	twilioToken := getEnv("TWILIO_AUTH_TOKEN")
	twilioFrom := getEnv("TWILIO_FROM_PHONE")

	return AppConfig{
		ServerPort:       httpPort,
		Dsn:              dsn,
		AppSecret:        appSecret,
		TwilioAccountSID: twilioSID,
		TwilioAuthToken:  twilioToken,
		TwilioFromPhone:  twilioFrom,
	}, nil
}

// getEnv trims and returns the environment variable, or returns error if missing
func getEnv(key string) string {
	val := strings.TrimSpace(os.Getenv(key))
	if val == "" {
		panic(fmt.Sprintf("missing required environment variable: %s", key))
	}
	return val
}
