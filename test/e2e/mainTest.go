package e2e

import (
	"fmt"
	"go-ecommerce-app/config"
	"go-ecommerce-app/internal/api"
	"go-ecommerce-app/internal/domain"
	"log"
	"os"
	"testing"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/gofiber/fiber/v2"
)

var (
	app *fiber.App
	db  *gorm.DB
)

func TestMain(m *testing.M) {
	fmt.Println("🔧 Loading test environment...")

	// Load test config
	cfg, err := config.SetupEnv("../../.env.test")
	if err != nil {
		log.Fatalf("❌ Failed to load test config: %v", err)
	}

	// Connect to test DB
	db, err = gorm.Open(postgres.Open(cfg.Dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ Failed to connect to test DB: %v", err)
	}

	// Reset and migrate schema
	if err := resetAndMigrateDB(); err != nil {
		log.Fatalf("❌ Failed to reset/migrate DB: %v", err)
	}

	// Setup the Fiber app
	app = api.CreateApp(cfg)

	// Run tests
	exitCode := m.Run()


	os.Exit(exitCode)
}

func resetAndMigrateDB() error {
	fmt.Println("🔄 Resetting test DB schema...")
	if err := db.Exec("DROP SCHEMA public CASCADE; CREATE SCHEMA public;").Error; err != nil {
		return fmt.Errorf("failed to reset schema: %w", err)
	}
	fmt.Println("🚀 Migrating schema...")
	return db.AutoMigrate(
		domain.User{},
	)
}
