package config

import (
	"log"
	"os"
)

type Config struct {
	DBDriver   string
	DBSource   string
	JWTSecret  string
	ServerPort string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

func LoadConfig() *Config {
	return &Config{
		// DBDriver:   "sqlite",
		// DBSource:   "rent-a-room.db",
		JWTSecret:  os.Getenv("JWT_SECRET"),
		ServerPort: os.Getenv("PORT"),
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASS"),
		DBName:     os.Getenv("DB_NAME"),
	}
}

func LoadEnvVars() {
	if err := os.Setenv("JWT_SECRET", "your-secret-key"); err != nil {
		log.Fatal("Failed to set JWT_SECRET")
	}
	if err := os.Setenv("PORT", "3000"); err != nil {
		log.Fatal("Failed to set PORT")
	}
}
