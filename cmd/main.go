package main

import (
	"LeaseEase/config"
	"LeaseEase/internal/handlers"
	"LeaseEase/internal/models"
	"LeaseEase/internal/repositories"
	"LeaseEase/internal/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// Load configuration
	config.LoadEnvVars()
	cfg := config.LoadConfig()

	// Initialize database
	db, err := gorm.Open(sqlite.Open(cfg.DBSource), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	// Migrate models
	db.AutoMigrate(&models.User{}, &models.Property{}, &models.Request{})

	// Initialize repositories, services, and handlers
	repositories := repositories.NewRepository(cfg, db)
	// userRepo := &repositories.UserRepository{db: db}
	services := services.NewService(repositories)
	// authService := &services.AuthService{UserRepo: userRepo}
	handlers := handlers.NewHandler(services)
	// authHandler := &handlers.AuthHandler{AuthService: authService}

	// Set up Fiber app
	app := fiber.New()

	// Auth routes
	app.Post("/auth/register", handlers.UserHandler.Register)

	// Start server
	app.Listen(":" + cfg.ServerPort)
}
