package storage

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// DeleteFile removes an image from S3/Supabase Storage
func DeleteFile(s3Client *s3.Client, bucketName, key string) error {
	input := &s3.DeleteObjectInput{
		Bucket: &bucketName,
		Key:    &key,
	}

	_, err := s3Client.DeleteObject(context.TODO(), input)
	if err != nil {
		return fmt.Errorf("failed to delete object: %v", err)
	}

	fmt.Println("🗑️ Deleted:", key)
	return nil
}
