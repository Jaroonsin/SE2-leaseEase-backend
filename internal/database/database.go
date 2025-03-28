package database

import (
	"LeaseEase/config"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
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

// InitS3Client initializes an S3-compatible client

func InitS3Client(cfg *config.Config) (*s3.Client, error) {

	provider := os.Getenv("S3_PROVIDER")

	if provider == "" {
		provider = "supabase"
	}

	var endpoint, region, accessKey, secretKey string

	if provider == "supabase" {
		endpoint = os.Getenv("SUPABASE_URL")
		region = os.Getenv("SUPABASE_REGION")
		accessKey = os.Getenv("SUPABASE_ACCESS_KEY")
		secretKey = os.Getenv("SUPABASE_SECRET_KEY")
		log.Printf("S3 provider: %s, endpoint: %s, region: %s", provider, endpoint, region)

	} else {

		region = os.Getenv("AWS_REGION")
		accessKey = os.Getenv("AWS_ACCESS_KEY_ID")
		secretKey = os.Getenv("AWS_SECRET_ACCESS_KEY")

	}

	if accessKey == "" || secretKey == "" {
		log.Fatal("Access key and secret key must be set")
	}

	var AWScfg aws.Config
	var err error

	AWScfg, err = awsconfig.LoadDefaultConfig(context.TODO(),
		awsconfig.WithRegion(region),
		awsconfig.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKey, secretKey, "")),
	)

	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	if provider == "supabase" {
		// Create S3 client with custom endpoint options for Supabase
		s3Client := s3.NewFromConfig(AWScfg, func(o *s3.Options) {
			o.BaseEndpoint = aws.String(endpoint)
			o.UsePathStyle = true

		})
		log.Println("S3 supabase client initialized successfully.")
		return s3Client, nil

	}
	log.Println("S3 aws client initialized successfully.")
	// Default AWS S3 client
	return s3.NewFromConfig(AWScfg), nil
}
