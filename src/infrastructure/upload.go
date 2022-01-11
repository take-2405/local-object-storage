package infrastructure

import (
	"github.com/minio/minio-go"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

func (m minioMethods) UploadImage(file *multipart.FileHeader, fileName string) error {
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
	info, err := m.Client.FPutObject(ctx, "aaa", objectName, filePath, minio.PutObjectOptions{ContentType: fileType})
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println(info)
	return err
}
