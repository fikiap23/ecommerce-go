package api

import (
	"go-ecommerce-app/config"
	"go-ecommerce-app/internal/api/rest"
	"go-ecommerce-app/internal/api/rest/handlers"
	"go-ecommerce-app/internal/domain"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func StartServer(config config.AppConfig) {
	app := fiber.New()
	log.Println("Server started on", config.ServerPort)

	db, err := gorm.Open(postgres.Open(config.Dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("failed to connect database, error: ", err)
	}

	log.Println("Database connected")

	// run migrations
	db.AutoMigrate(&domain.User{})

	rh := &rest.RestHandler{
		App: app,
		DB: db,
	}

	setupRoutes(rh)

	app.Listen(config.ServerPort)
}


func setupRoutes(rh *rest.RestHandler) {
	// setup routes
	handlers.SetupUserRoutes(rh)
}