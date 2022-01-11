package application

import (
	"local-object-storage/src/domain/repository"
	"log"
	"mime/multipart"
)

type UploadUseCase interface {
	UploadImage(file *multipart.FileHeader, fileName string, bucketName string) error
}

type uploadUseCase struct {
	minio repository.MinioRepository
}

func NewUploadUseCase(minio repository.MinioRepository) UploadUseCase {
	return &uploadUseCase{minio: minio}
}

func (uu uploadUseCase) UploadImage(file *multipart.FileHeader, fileName string, bucketName string) error {
	buckets, err := uu.minio.GetBucketLists()
	if err != nil {
		log.Println(err)
		return err
	}
	for i, j := range buckets {
		if j == bucketName {
			break
		}
		if i == len(buckets)-1 {
			return err
		}
	}
	err = uu.minio.UploadImage(file, fileName, bucketName)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
