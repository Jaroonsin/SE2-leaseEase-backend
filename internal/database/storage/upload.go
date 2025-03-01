package storage

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// UploadFile uploads an image to Supabase Storage or Amazon S3
func UploadFile(s3Client *s3.Client, bucketName, filePath, key string) error {
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
		Bucket:        &bucketName,
		Key:           &key,
		Body:          file,
		ContentLength: aws.Int64(fileInfo.Size()),
		ContentType:   aws.String("image/jpeg"),
	}

	_, err = s3Client.PutObject(context.TODO(), input)
	if err != nil {
		return fmt.Errorf("failed to upload object: %v", err)
	}

	log.Println("✅ Uploaded:", key)
	return nil
}
