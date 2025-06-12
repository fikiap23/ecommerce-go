package api

import (
	"go-ecommerce-app/config"
	"go-ecommerce-app/internal/api/rest"
	"go-ecommerce-app/internal/api/rest/handlers"
	"go-ecommerce-app/internal/domain"
	"go-ecommerce-app/internal/helper"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func CreateApp(config config.AppConfig) *fiber.App {
	app := fiber.New()

	db, err := gorm.Open(postgres.Open(config.Dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database, error: ", err)
	}

	db.AutoMigrate(&domain.User{})
	auth := helper.SetupAuth(config.AppSecret)

	rh := &rest.RestHandler{
		App: app,
		DB: db,
		Auth: auth,
	}

	setupRoutes(rh)
	return app
}


func StartServer(config config.AppConfig) {
	app := CreateApp(config)
	log.Println("Server started on", config.ServerPort)
	app.Listen(config.ServerPort)
}



func setupRoutes(rh *rest.RestHandler) {
	// setup routes
	handlers.SetupUserRoutes(rh)
}