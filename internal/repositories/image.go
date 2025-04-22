package repositories

import (
	"context"
	"mime/multipart"
)

type ImageRepository interface {
	UploadFile(ctx context.Context, key string, fileHeader *multipart.FileHeader) (string, error)
	// DownloadFile(bucket, key, destination string) error
	// DeleteFile(bucket, key string) error
	// UpdateFile(bucket, filePath, key string) error
}
