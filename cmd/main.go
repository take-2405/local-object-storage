package main

import (
    "log"

    "github.com/minio/minio-go/v7"
    "github.com/minio/minio-go/v7/pkg/credentials"
)

func main() {
    endpoint := "localhost:9090"
    accessKeyID := "minio"
    secretAccessKey := "minio123"
    useSSL := false

    // Initialize minio client sobject.
    minioClient, err := minio.New(endpoint, &minio.Options{
        Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
        Secure: useSSL,
    })
    if err != nil {
        log.Fatalln(err)
    }

    log.Printf("%#v\n", minioClient) // minioClient is now set up
}