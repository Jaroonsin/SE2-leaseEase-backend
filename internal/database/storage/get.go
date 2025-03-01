package storage

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// DownloadFile retrieves a file from S3/Supabase Storage
func DownloadFile(s3Client *s3.Client, bucketName, key, destination string) error {
	input := &s3.GetObjectInput{
		Bucket: &bucketName,
		Key:    &key,
	}

	output, err := s3Client.GetObject(context.TODO(), input)
	if err != nil {
		return fmt.Errorf("failed to retrieve object: %v", err)
	}
	defer output.Body.Close()

	// Create destination file
	destFile, err := os.Create(filepath.Join(destination, filepath.Base(key)))
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer destFile.Close()

	// Copy content
	_, err = io.Copy(destFile, output.Body)
	if err != nil {
		return fmt.Errorf("failed to write file: %v", err)
	}

	fmt.Println("✅ Downloaded:", key)
	return nil
}
