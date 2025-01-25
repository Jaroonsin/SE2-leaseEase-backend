package main

import (
	"LeaseEase/config"
	"LeaseEase/internal/handlers"
	"LeaseEase/internal/models"
	"LeaseEase/internal/repositories"
	"LeaseEase/internal/services"
	"log"

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
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// AutoMigrate models
	err = db.AutoMigrate(
		&models.User{},
		&models.Customer{},
		&models.Admin{},
		&models.Lessor{},
		&models.Lessee{},
		&models.PremiumLessor{},
		&models.Advertisement{},
		&models.MarketSlot{},
		&models.Request{},
		&models.Transaction{},
		&models.Review{},
		&models.LessorReview{},
		&models.SlotReview{},
		&models.Problem{},
		&models.ProblemTag{},
		&models.Solve{},
		&models.ChatMessage{},
		&models.Report{},
	)
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	log.Println("Database migration completed.")

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
