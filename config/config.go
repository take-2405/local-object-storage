package config

import (
	"log"
	"github.com/kelseyhightower/envconfig"
)

type minioConfig struct{
	EndPoint string `envconfig:"DB_USER" default:"localhost:9090"`
	AccessKey string `envconfig:"DB_PASSWORD" default:"minio"`
	SecretAccessKey string   `envconfig:"DB_IP" default:"minio123"`
}

func GetMinioConfig()(string,string,string){
	var config minioConfig
	if err := envconfig.Process("",&config);err !=nil{
		log.Fatal("Unable to get Minio config")
	}
	return config.EndPoint ,config.AccessKey,config.SecretAccessKey;
}