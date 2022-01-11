package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"local-object-storage/src/application"
	"local-object-storage/src/infrastructure"
	"local-object-storage/src/presentation"
	"local-object-storage/src/presentation/handler"
	"log"
	"os"
	"time"
)

func main() {
	minio := infrastructure.NewMinioRepository()
	bucketUseCase := application.NewBucketUseCase(minio)
	uploadUseCase := application.NewUploadUseCase(minio)

	bucketHandler := handler.NewBucketHandler(bucketUseCase)
	uploadHandler := handler.NewUploadHandler(uploadUseCase)

	s := gin.Default()
	s.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "accept", "origin", "Cache-Control", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		MaxAge: 15 * time.Second,
	}))

	presentation.InitRouting(s, bucketHandler, uploadHandler)

	if err := s.Run(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

}
