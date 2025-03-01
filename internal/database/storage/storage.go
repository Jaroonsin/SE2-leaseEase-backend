package storage

import "github.com/aws/aws-sdk-go-v2/service/s3"

// StorageProvider defines the interface for storage services
type Storage interface {
	UploadFile(filePath, key string) error
	DownloadFile(key, destination string) error
	DeleteFile(key string) error
	UpdateFile(filePath, key string) error
}

type storage struct {
	client     *s3.Client
	bucketname string
}

// NewStorage returns a new storage service
func NewStorage() Storage {
	client, bucketname, err := InitS3Client()
	if err != nil {
		panic(err)
	}

	return &storage{
		client:     client,
		bucketname: bucketname,
	}
}

// UploadFile uploads an image to Supabase Storage or Amazon S3
func (s *storage) UploadFile(filePath, key string) error {
	return UploadFile(s.client, s.bucketname, filePath, key)
}

// DownloadFile retrieves a file from S3/Supabase Storage
func (s *storage) DownloadFile(key, destination string) error {
	return DownloadFile(s.client, s.bucketname, key, destination)
}

// DeleteFile removes an image from S3/Supabase Storage
func (s *storage) DeleteFile(key string) error {
	return DeleteFile(s.client, s.bucketname, key)
}

// UpdateFile replaces an existing file in S3/Supabase Storage
func (s *storage) UpdateFile(filePath, key string) error {
	return UpdateFile(s.client, s.bucketname, filePath, key)
}
