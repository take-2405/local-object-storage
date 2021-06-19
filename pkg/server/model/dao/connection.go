package dao

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"local-object-storage/config"
	"log"
)

type MinioConfig struct{
	EndPoint string
	AccessKey string
	SecretAccessKey string
}


type Minio struct{
	Client *minio.Client
}

func New() Minio {
	var Config MinioConfig
	var Minio Minio
	var err error
	Config.EndPoint,Config.AccessKey,Config.SecretAccessKey=config.GetMinioConfig()
	Minio.Client, err = minio.New(Config.EndPoint, &minio.Options{
		Creds:  credentials.NewStaticV4(Config.AccessKey, Config.SecretAccessKey, ""),
		Secure: false,
	})
	if err != nil {
		log.Fatalln(err)
	}
	return Minio
}