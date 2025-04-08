package repositories

import (
	"context"
	"fmt"
	"log"
	"mime/multipart"
	"path/filepath"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type imageRepository struct {
	s3         *s3.Client
	bucketName string
}

// NewImageRepository returns a new image repository
func NewImageRepository(s3 *s3.Client) ImageRepository {
	return &imageRepository{
		s3:         s3,
		bucketName: "images",
	}
}

func (r *imageRepository) UploadFile(ctx context.Context, key string, fileHeader *multipart.FileHeader) (string, error) {
	file, err := fileHeader.Open()
	if err != nil {
		return "", fmt.Errorf("failed to open file '%s': %w", fileHeader.Filename, err)
	}

	// Determine Content-Type based on file extension (more robust)
	contentType := getContentType(fileHeader.Filename)

	ext := filepath.Ext(fileHeader.Filename)

	key = fmt.Sprintf("%s%s", key, ext)
	// Upload request
	input := &s3.PutObjectInput{
		Bucket:        aws.String(r.bucketName),
		Key:           aws.String(key),
		Body:          file,
		ContentType:   aws.String(contentType),
		ContentLength: aws.Int64(fileHeader.Size),
	}

	_, err = r.s3.PutObject(context.TODO(), input)
	if err != nil {
		return "", fmt.Errorf("failed to upload object '%s' to bucket '%s': %w", key, r.bucketName, err)
	}

	s3URL := fmt.Sprintf("https://khgvvndokfqibyevapec.supabase.co/storage/v1/object/public/%s/%s", r.bucketName, key)
	log.Printf("✅ Uploaded '%s' to '%s/%s'", fileHeader.Filename, r.bucketName, key)
	file.Close()
	return s3URL, nil
}

func getContentType(filename string) string {
	ext := filepath.Ext(filename)
	switch ext {
	case ".jpeg", ".jpg":
		return "image/jpeg"
	case ".png":
		return "image/png"
	case ".gif":
		return "image/gif"
	case ".webp":
		return "image/webp"
	default:
		return "application/octet-stream" // Default binary type
	}
}

// // ImageRepository defines the interface for image storage
// func (r *imageRepository) UploadFile(bucket, filePath, key string) (string, error) {
// 	file, err := os.Open(filePath)
// 	if err != nil {
// 		return "", fmt.Errorf("failed to open file: %v", err)
// 	}
// 	defer file.Close()

// 	// Get file info
// 	fileInfo, err := file.Stat()
// 	if err != nil {
// 		return "", fmt.Errorf("failed to get file info: %v", err)
// 	}

// 	// Upload request
// 	input := &s3.PutObjectInput{
// 		Bucket:        &bucket,
// 		Key:           &key,
// 		Body:          file,
// 		ContentLength: aws.Int64(fileInfo.Size()),
// 		ContentType:   aws.String("image/jpeg"),
// 	}

// 	_, err = r.s3.PutObject(context.TODO(), input)
// 	if err != nil {
// 		return "", fmt.Errorf("failed to upload object: %v", err)
// 	}

// 	log.Println("✅ Uploaded:", key)
// 	return key, nil
// }

// // DownloadFile retrieves a file from S3/Supabase Storage
// func (r *imageRepository) UpdateFile(bucket, filePath, key string) error {
// 	fmt.Println("🔄 Updating:", key)

// 	// Delete old file
// 	err := r.DeleteFile(bucket, key)
// 	if err != nil {
// 		return fmt.Errorf("failed to delete old file: %v", err)
// 	}

// 	// Upload new file
// 	key, err = r.UploadFile(bucket, filePath, key)
// 	if err != nil {
// 		return fmt.Errorf("failed to upload new file: %v", err)
// 	}

// 	fmt.Println("✅ Updated:", key)
// 	return nil
// }

// // DeleteFile removes an image from S3/Supabase Storage
// func (r *imageRepository) DeleteFile(bucket, key string) error {
// 	input := &s3.DeleteObjectInput{
// 		Bucket: &bucket,
// 		Key:    &key,
// 	}

// 	_, err := r.s3.DeleteObject(context.TODO(), input)
// 	if err != nil {
// 		return fmt.Errorf("failed to delete object: %v", err)
// 	}

// 	fmt.Println("🗑️ Deleted:", key)
// 	return nil
// }

// // DownloadFile retrieves a file from S3/Supabase Storage
// func (r *imageRepository) DownloadFile(bucket, key, destination string) error {
// 	input := &s3.GetObjectInput{
// 		Bucket: &bucket,
// 		Key:    &key,
// 	}

// 	output, err := r.s3.GetObject(context.TODO(), input)
// 	if err != nil {
// 		return fmt.Errorf("failed to retrieve object: %v", err)
// 	}
// 	defer output.Body.Close()

// 	// Create destination file
// 	destFile, err := os.Create(filepath.Join(destination, filepath.Base(key)))
// 	if err != nil {
// 		return fmt.Errorf("failed to create file: %v", err)
// 	}
// 	defer destFile.Close()

// 	// Copy content
// 	_, err = io.Copy(destFile, output.Body)
// 	if err != nil {
// 		return fmt.Errorf("failed to write file: %v", err)
// 	}

// 	fmt.Println("✅ Downloaded:", key)
// 	return nil
// }
