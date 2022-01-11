package infrastructure

import (
	"context"
	"github.com/minio/minio-go/v7"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

func (m *minioRepository) UploadImage(file *multipart.FileHeader, fileName string, bucket string) error {
	var err error
	ctx := context.Background()
	f, err := file.Open()
	if err != nil {
		return err
	}
	defer f.Close()
	size := file.Size
	buffer := make([]byte, size)

	newFile, err := os.Create("./../images/" + fileName + ".png")
	defer newFile.Close()
	io.Copy(newFile, f)

	fileType := http.DetectContentType(buffer)
	objectName := fileName + ".png"
	filePath := "./../images/" + fileName + ".png"
	info, err := m.Client.FPutObject(ctx, bucket, objectName, filePath, minio.PutObjectOptions{ContentType: fileType})
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println(info)
	return err
}
