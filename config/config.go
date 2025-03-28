package config

import (
	"os"
)

type Config struct {
	// Client settings
	ClientURL string

	// Server settings
	ServerEnv  string
	ServerPort string

	// Database settings: PGSQL
	DBURL      string
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

	// S3 Provider: "supabase" or "amazon"
	S3Provider string

	// S3 supabase settings
	SupabaseRegion    string
	SupabaseURL       string
	SupabaseAccessKey string
	SupabaseSecretKey string

	// S3 amazon settings
	AmazonRegion    string
	AmazonAccessKey string
	AmazonSecretKey string
}

func LoadConfig() *Config {
	return &Config{
		// Client settings
		ClientURL: os.Getenv("CLIENT_URL"),

		// Server settings
		ServerEnv:  os.Getenv("SERVER_ENV"),
		ServerPort: os.Getenv("SERVER_PORT"),

		// Database: PGSQL
		DBURL:      os.Getenv("DB_URL"),
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

		// S3 Provider: "supabase" or "amazon"
		S3Provider: os.Getenv("S3_PROVIDER"),
		// S3 supabase settings
		SupabaseRegion:    os.Getenv("SUPABASE_REGION"),
		SupabaseURL:       os.Getenv("SUPABASE_URL"),
		SupabaseAccessKey: os.Getenv("SUPABASE_ACCESS_KEY"),
		SupabaseSecretKey: os.Getenv("SUPABASE_SECRET_KEY"),

		// S3 amazon settings
		AmazonRegion:    os.Getenv("AWS_REGION"),
		AmazonAccessKey: os.Getenv("AWS_ACCESS_KEY_ID"),
		AmazonSecretKey: os.Getenv("AWS_SECRET_ACCESS_KEY"),
	}
}

func LoadEnv() string {
	return os.Getenv("SERVER_ENV")
}
