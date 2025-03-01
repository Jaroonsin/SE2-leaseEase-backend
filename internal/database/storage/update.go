package storage

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// UpdateFile replaces an existing file in S3/Supabase Storage
func UpdateFile(s3Client *s3.Client, bucketName, filePath, key string) error {
	fmt.Println("🔄 Updating:", key)

	// Delete old file
	err := DeleteFile(s3Client, bucketName, key)
	if err != nil {
		return fmt.Errorf("failed to delete old file: %v", err)
	}

	// Upload new file
	err = UploadFile(s3Client, bucketName, filePath, key)
	if err != nil {
		return fmt.Errorf("failed to upload new file: %v", err)
	}

	fmt.Println("✅ Updated:", key)
	return nil
}
