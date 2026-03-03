package main

import (
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

	host := os.Getenv("APP_HOST")
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	address := host + ":" + port

	log.Printf("Starting server on %s...", address)
	err = app.Listen(address)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
