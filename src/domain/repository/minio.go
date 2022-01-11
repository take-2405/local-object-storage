package domain

import "mime/multipart"

type MinioRepository interface {
	GetBucketLists() ([]string, error)
	UploadImage(file *multipart.FileHeader, fileName string) error
	CreateBucket(bucketName string) error
}
