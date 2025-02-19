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
		&models.Property{},
		&models.Reservation{},
		&models.Review{},
		&models.LessorReview{},
		&models.PropertyReview{},
	)
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	log.Println("Database migration completed.")
}
