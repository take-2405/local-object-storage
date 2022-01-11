package infrastructure

import (
	"errors"
	"github.com/minio/minio-go"
	"io/ioutil"
	"log"
)

type minioMethods struct {
	Client *minio.Client
}

func newMinioClient(client *minio.Client) methods {
	return &minioMethods{Client: client}
}

func (m minioMethods) GetBucketLists() ([]string, error) {
	var list string
	var lists []string
	raws, err := ioutil.ReadDir("./../docker/minio/data")
	if err != nil {
		return lists, err
	}
	for i := 1; i < len(raws); i++ {
		list = raws[i].Name()
		lists = append(lists, list)
	}
	return lists, nil
}

func (m minioMethods) CreateBucket(bucketName string) error {
	var err error
	ctx := context.Background()
	location := "us-east-1"
	err = m.Client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
	if err != nil {
		exists, errBucketExists := m.Client.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			return errors.New(bucketName + "is already exist")
		} else {
			return err
		}
	} else {
		log.Printf("Successfully created %s\n", bucketName)
	}
	return err
}
