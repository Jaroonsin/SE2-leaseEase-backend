package database

import (
	"LeaseEase/config"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(cfg *config.Config) (*gorm.DB, error) {
	var dsn string

	if cfg.DBURL != "" {
		dsn = cfg.DBURL
	} else {
		dsn = fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Bangkok",
			cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName,
		)
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	log.Println("Database connection successfully.")

	// Run migrations
	RunMigrations(db)
	// Setup PostgreSQL function
	SetupFunctions(db)
	// Setup PostgreSQL trigger
	SetupTriggers(db)

	return db, nil
}
