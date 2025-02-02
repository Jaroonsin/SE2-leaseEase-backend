package database

import (
	"LeaseEase/internal/models"
	"log"

	"gorm.io/gorm"
)
func RunMigrations(db *gorm.DB) {

	// AutoMigrate models
	err := db.AutoMigrate(
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
}