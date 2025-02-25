package database

import (
	"log"
	"os"

	"gorm.io/gorm"
)

// SetupTrigger loads and executes SQL trigger file
func SetupTriggers(db *gorm.DB) error {
	log.Println("Setting up PostgreSQL trigger...")

	sqlFile := "internal/database/sql/trigger_update_property.sql"
	sql, err := os.ReadFile(sqlFile)
	if err != nil {
		log.Println("Error reading SQL file:", err)
		return err
	}

	if err := db.Exec(string(sql)).Error; err != nil {
		log.Println("Error executing trigger setup:", err)
		return err
	}

	log.Println("PostgreSQL trigger created successfully.")
	return nil
}
