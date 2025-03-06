package repositories

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type imageRepository struct {
	s3 *s3.Client
}

// NewImageRepository returns a new image repository
func NewImageRepository(s3 *s3.Client) ImageRepository {
	return &imageRepository{
		s3: s3,
	}
}

// ImageRepository defines the interface for image storage
func (r *imageRepository) UploadFile(bucket, filePath, key string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	// Get file info
	fileInfo, err := file.Stat()
	if err != nil {
		return fmt.Errorf("failed to get file info: %v", err)
	}

	// Upload request
	input := &s3.PutObjectInput{
		Bucket:        &bucket,
		Key:           &key,
		Body:          file,
		ContentLength: aws.Int64(fileInfo.Size()),
		ContentType:   aws.String("image/jpeg"),
	}

	_, err = r.s3.PutObject(context.TODO(), input)
	if err != nil {
		return fmt.Errorf("failed to upload object: %v", err)
	}

	log.Println("✅ Uploaded:", key)
	return nil
}

// DownloadFile retrieves a file from S3/Supabase Storage
func (r *imageRepository) UpdateFile(bucket, filePath, key string) error {
	fmt.Println("🔄 Updating:", key)

	// Delete old file
	err := r.DeleteFile(bucket, key)
	if err != nil {
		return fmt.Errorf("failed to delete old file: %v", err)
	}

	// Upload new file
	err = r.UploadFile(bucket, filePath, key)
	if err != nil {
		return fmt.Errorf("failed to upload new file: %v", err)
	}

	fmt.Println("✅ Updated:", key)
	return nil
}

// DeleteFile removes an image from S3/Supabase Storage
func (r *imageRepository) DeleteFile(bucket, key string) error {
	input := &s3.DeleteObjectInput{
		Bucket: &bucket,
		Key:    &key,
	}

	_, err := r.s3.DeleteObject(context.TODO(), input)
	if err != nil {
		return fmt.Errorf("failed to delete object: %v", err)
	}

	fmt.Println("🗑️ Deleted:", key)
	return nil
}

// DownloadFile retrieves a file from S3/Supabase Storage
func (r *imageRepository) DownloadFile(bucket, key, destination string) error {
	input := &s3.GetObjectInput{
		Bucket: &bucket,
		Key:    &key,
	}

	output, err := r.s3.GetObject(context.TODO(), input)
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
