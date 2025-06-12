package main

import (
	"go-ecommerce-app/config"
	"go-ecommerce-app/internal/api"
	"go-ecommerce-app/pkg/utils"

	"github.com/gofiber/fiber/v2/log"
)

func main() {
	
	cfg, err := config.SetupEnv(".env.test")
	if err != nil {
		log.Fatalf("config error: %v", err)
		panic(err)
	}

	if err := utils.LoadValidationMessages(); err != nil {
		log.Fatalf("Error loading i18n validation messages: %v", err)
	}
	api.StartServer(cfg)
}