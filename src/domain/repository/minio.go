package repository

import "mime/multipart"

type MinioRepository interface {
	GetBucketLists() ([]string, error)
	UploadImage(file *multipart.FileHeader, fileName string, bucket string) error
	CreateBucket(bucketName string) error
}
