package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	ServerPort string
	Dsn string
}

func SetupEnv()(cfg AppConfig, err error) {
	if os.Getenv("APP_ENV") == "dev" {
		godotenv.Load()
	}
	htppPort := os.Getenv("HTTP_PORT")
	if len(htppPort) < 1 {
		return AppConfig{}, errors.New("HTTP_PORT is not set")
	}

	Dsn := os.Getenv("DSN")
	if len(Dsn) < 1 {
		return AppConfig{}, errors.New("DSN is not set")
	}

	return AppConfig{
		ServerPort: htppPort,
		Dsn: Dsn,
	}, nil
}