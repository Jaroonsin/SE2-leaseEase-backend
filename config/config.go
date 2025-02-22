package config

import (
	"os"
)

type Config struct {
	// Server settings
	// ServerName string
	ServerEnv string
	// ServerURL  string
	// ServerHost string
	ServerPort string

	// Database settings: PGSQL
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	// DBSSLMode  string
	// DBTimeZone string

	// JWT settings
	JWTApiSecret string
	// JWTAccessTokenSecret string
	// JWTRefreshTokenSecret string
	// JWTAccessTokenExpiration string
	// JWTRefreshTokenExpiration string

	// Email settings
	EmailHost     string
	EmailPort     string
	EmailUser     string
	EmailPassword string

	// Website settings
	ResetPassURL string

	// Omise settings
	OmisePublicKey string
	OmiseSecretKey string
}

func LoadConfig() *Config {
	return &Config{
		// Server settings
		// ServerName: os.Getenv("SERVER_NAME"),	//not used
		ServerEnv: os.Getenv("SERVER_ENV"),
		// ServerURL:  os.Getenv("SERVER_URL"), 	//not used
		// ServerHost: os.Getenv("SERVER_HOST"),	//not used
		ServerPort: os.Getenv("SERVER_PORT"),

		// Database: PGSQL
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASS"),
		DBName:     os.Getenv("DB_NAME"),
		// DBSSLMode:  os.Getenv("DB_SSL_MODE"),	//not used
		// DBTimeZone: os.Getenv("DB_TIMEZONE"), 	//not used

		// # JWT settings
		JWTApiSecret: os.Getenv("JWT_API_SECRET_KEY"),
		// JWTAccessTokenSecret: os.Getenv("JWT_ACCESS_TOKEN_SECRET"),				//not used
		// JWTRefreshTokenSecret: os.Getenv("JWT_REFRESH_TOKEN_SECRET"),			//not used
		// JWTAccessTokenExpiration: os.Getenv("JWT_ACCESS_TOKEN_EXPIRATION"),		//not used
		// JWTRefreshTokenExpiration: os.Getenv("JWT_REFRESH_TOKEN_EXPIRATION"),	//not used

		// Email settings
		EmailHost:     os.Getenv("EMAIL_HOST"),
		EmailPort:     os.Getenv("EMAIL_PORT"),
		EmailUser:     os.Getenv("EMAIL_USER"),
		EmailPassword: os.Getenv("EMAIL_PASS"),

		// Website settings
		ResetPassURL: os.Getenv("RESET_PASSWORD_URL"),

		// Omise settings
		OmisePublicKey: os.Getenv("OMISE_PUBLIC_KEY"),
		OmiseSecretKey: os.Getenv("OMISE_SECRET_KEY"),
	}
}

func LoadEnv() string {
	return os.Getenv("SERVER_ENV")
}
