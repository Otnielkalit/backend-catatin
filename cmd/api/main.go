package main

import (
	"fmt"
	"log"
	"os"

	"be-catatin/config"
	"be-catatin/internal/controller"
	"be-catatin/internal/entity"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.InitEnv()
	config.InitDB()

	// Auto Migrate Database
	log.Println("Running Database Migration...")
	err := config.DB.AutoMigrate(
		&entity.User{},
		&entity.Budget{},
		&entity.Category{},
		&entity.Expense{},
	)
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
	log.Println("Database Migration completed successfully!")

	app := fiber.New()

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	// Setup Routes
	controller.SetupRoutes(app, config.DB)

	// Determine Port
	port := os.Getenv("PORT")
	if port == "" {
		port = os.Getenv("APP_PORT")
		if port == "" {
			port = "8080" // Default fallback
		}
	}
	host := os.Getenv("APP_HOST")
	if host == "" {
		host = "0.0.0.0" // Ensure it listens on all interfaces (crucial for Cloud deployments)
	}

	address := fmt.Sprintf("%s:%s", host, port)
	log.Printf("Starting server on %s...\n", address)

	err = app.Listen(address)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
