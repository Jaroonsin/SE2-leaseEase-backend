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
		&models.Payment{},
	)

	sqlFunction := `
	CREATE OR REPLACE FUNCTION update_property_rating()
	RETURNS TRIGGER AS $$
	BEGIN
	    UPDATE properties
	    SET avg_rating = (
	        SELECT COALESCE(AVG(reviews.rating), 0)
	        FROM property_reviews
	        JOIN reviews ON property_reviews.review_id = reviews.id
	        WHERE property_reviews.property_id = NEW.property_id
	    )
	    WHERE id = NEW.property_id;
	    
	    RETURN NEW;
	END;
	$$ LANGUAGE plpgsql;

	CREATE OR REPLACE TRIGGER trigger_update_property_rating
	AFTER INSERT ON property_reviews
	FOR EACH ROW
	EXECUTE FUNCTION update_property_rating();
	`

	// Execute SQL commands
	db.Exec(sqlFunction)

	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	log.Println("Database migration completed.")
}
