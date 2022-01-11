package infrastructure

import (
	"github.com/minio/minio-go"
	"github.com/minio/minio-go/pkg/credentials"
	"local-object-storage/config"
	"local-object-storage/src/domain/repository"
	"log"
)

type MinioConfig struct {
	EndPoint        string
	AccessKey       string
	SecretAccessKey string
}

type Minio struct {
	Client  *minio.Client
	Methods Methods
}

type Methods struct {
	MinioMethods repository.Minio
}

//type methods interface {
//	GetBucketLists() ([]string, error)
//	UploadImage(file *multipart.FileHeader, fileName string) error
//	CreateBucket(bucketName string) error
//}

func New() Minio {
	var Config MinioConfig
	var Minio Minio
	var err error
	Config.EndPoint, Config.AccessKey, Config.SecretAccessKey = config.GetMinioConfig()
	Minio.Client, err = minio.New(Config.EndPoint, &minio.Options{
		Creds:  credentials.NewStaticV4(Config.AccessKey, Config.SecretAccessKey, ""),
		Secure: false,
	})
	if err != nil {
		log.Fatalln(err)
	}

	Minio.Methods.MinioMethods = newMinioClient(Minio.Client)
	return Minio
}
