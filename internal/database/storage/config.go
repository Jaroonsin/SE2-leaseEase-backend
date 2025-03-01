package storage

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// InitS3Client initializes an S3-compatible client
func InitS3Client() (*s3.Client, string, error) {
	provider := os.Getenv("S3_PROVIDER")
	if provider == "" {
		provider = "supabase"
	}

	var endpoint, region, accessKey, secretKey, bucketName string

	if provider == "supabase" {
		endpoint = os.Getenv("SUPABASE_URL")
		region = os.Getenv("SUPABASE_REGION")
		accessKey = os.Getenv("SUPABASE_ACCESS_KEY")
		secretKey = os.Getenv("SUPABASE_SECRET_KEY")
		bucketName = "lessor-images"
	} else {
		region = os.Getenv("AWS_REGION")
		accessKey = os.Getenv("AWS_ACCESS_KEY_ID")
		secretKey = os.Getenv("AWS_SECRET_ACCESS_KEY")
		bucketName = "lessor-images"
	}

	if accessKey == "" || secretKey == "" {
		log.Fatal("Access key and secret key must be set")
	}

	var cfg aws.Config
	var err error

	cfg, err = config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(region),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKey, secretKey, "")),
	)

	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	if provider == "supabase" {
		// Create S3 client with custom endpoint options for Supabase
		s3Client := s3.NewFromConfig(cfg, func(o *s3.Options) {
			o.BaseEndpoint = aws.String(endpoint)
			o.UsePathStyle = true
		})
		return s3Client, bucketName, nil
	}

	// Default AWS S3 client
	return s3.NewFromConfig(cfg), bucketName, nil
}
