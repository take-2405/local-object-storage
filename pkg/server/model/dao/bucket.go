package dao

import (
	"errors"
	"context"
	"github.com/minio/minio-go/v7"
	"io/ioutil"
	"log"
)

func CreateBuckt(minioClient *minio.Client,bucketName string)error{
	var err error
	ctx := context.Background()
	location := "us-east-1"
	err = minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
	if err != nil {
		exists, errBucketExists := minioClient.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			return errors.New(bucketName+"is already exist")
		} else {
			return err
		}
	} else {
		log.Printf("Successfully created %s\n", bucketName)
	}
	return err
}

func BucketList() ([]string, error) {
	var list string
	var lists []string
	raws, err := ioutil.ReadDir("./../docker/minio/data")
	if err != nil {
		return lists, err
	}
	for i:=1;i< len(raws);i++{
		list = raws[i].Name()
		lists=append(lists,list)
	}
	return lists, nil
}
