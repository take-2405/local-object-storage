package application

import (
	"local-object-storage/src/domain/repository"
	"log"
)

type BucketUseCase interface {
	CreateBucket(name string) error
	GetBucketLists() ([]string, error)
}

type bucketUseCase struct {
	minio repository.MinioRepository
}

func NewBucketUseCase(minio repository.MinioRepository) BucketUseCase {
	return &bucketUseCase{minio: minio}
}

func (bu bucketUseCase) CreateBucket(name string) error {
	if err := bu.minio.CreateBucket(name); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (bu bucketUseCase) GetBucketLists() ([]string, error) {
	buckets, err := bu.minio.GetBucketLists()
	if err != nil {
		log.Println(err)
		return buckets, err
	}
	return buckets, err
}
