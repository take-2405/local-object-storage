package infrastructure

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"local-object-storage/config"
	"local-object-storage/src/domain/repository"
	"log"
)

type minioConfig struct {
	EndPoint        string
	AccessKey       string
	SecretAccessKey string
}

type minioRepository struct {
	Client *minio.Client
}

func NewMinioRepository() repository.MinioRepository {
	var Config minioConfig
	var err error
	Config.EndPoint, Config.AccessKey, Config.SecretAccessKey = config.GetMinioConfig()
	conn, err := minio.New(Config.EndPoint, &minio.Options{
		Creds:  credentials.NewStaticV4(Config.AccessKey, Config.SecretAccessKey, ""),
		Secure: false,
	})
	if err != nil {
		log.Fatalln(err)
	}
	return &minioRepository{Client: conn}
}
