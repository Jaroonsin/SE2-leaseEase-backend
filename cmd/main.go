package main

import (
	"LeaseEase/cmd/server"
	"LeaseEase/config"
	"LeaseEase/internal/database"
	"LeaseEase/internal/handlers"
	"LeaseEase/internal/logs"
	"LeaseEase/internal/repositories"
	"LeaseEase/internal/services"
	"log"

	"github.com/joho/godotenv"
)

// @title LeaseEase API
// @version 2.0
// @description API documentation for LeaseEase.

// @contact.name API Support
// @contact.url http://www.example.com/support
// @contact.email support@example.com
// @host localhost:5000/api/v2
// @BasePath /
func main() {
	// Load configuration
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	cfg := config.LoadConfig()
	logger, err := logs.NewLogger()
	if err != nil {
		log.Printf("Failed to create logger: %v", err)
	}

	// Initialize database
	db, err := database.ConnectDB(cfg)
	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
	}

	// Initialize repositories, services, and handlers
	repositories := repositories.NewRepository(cfg, db)
	services := services.NewService(repositories, logger)
	handlers := handlers.NewHandler(services)

	servers := server.NewFiberHttpServer(cfg, logger, handlers)

	// Start server
	servers.Start()
}
