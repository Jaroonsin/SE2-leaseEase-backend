package database

import (
	"log"
	"os"

	"gorm.io/gorm"
)

// SetupFunction loads and executes SQL function file
func SetupFunctions(db *gorm.DB) error {
	log.Println("Setting up PostgreSQL function...")

	sqlFile := "internal/database/sql/update_property_rating.sql"
	sql, err := os.ReadFile(sqlFile)
	if err != nil {
		log.Println("Error reading SQL file:", err)
		return err
	}

	if err := db.Exec(string(sql)).Error; err != nil {
		log.Println("Error executing function setup:", err)
		return err
	}

	log.Println("PostgreSQL function created successfully.")
	return nil
}
