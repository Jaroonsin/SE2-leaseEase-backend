package repositories

type ImageRepository interface {
	UploadFile(bucket, filePath, key string) (string, error)
	DownloadFile(bucket, key, destination string) error
	DeleteFile(bucket, key string) error
	UpdateFile(bucket, filePath, key string) error
}
