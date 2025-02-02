package main

import (
	"LeaseEase/config"
	"LeaseEase/internal/database"
	"LeaseEase/internal/handlers"
	"LeaseEase/internal/repositories"
	"LeaseEase/internal/services"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	// Load configuration
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	config.LoadEnvVars()
	cfg := config.LoadConfig()

	// Initialize database
	db, err := database.ConnectDB(cfg)
	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
	}

	// Initialize repositories, services, and handlers
	repositories := repositories.NewRepository(cfg, db)
	services := services.NewService(repositories)
	handlers := handlers.NewHandler(services)

	// Set up Fiber app
	app := fiber.New()

	// Health check
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "success",
			"message": "Service is up and running!",
		})
	})
	// Auth routes
	app.Post("/auth/register", handlers.UserHandler.Register)

	// Start server
	app.Listen(":" + cfg.ServerPort)
}
