package dao

import(
	"context"
	"github.com/minio/minio-go/v7"
	"local-object-storage/pkg/decode"
	"local-object-storage/pkg/server/model/dto"
	"log"
)

func Upload(minioClient *minio.Client,request dto.UploadImageRequest)error{
	var err error
	var contentType string
	ctx := context.Background()
	if request.Type=="png" {
		contentType = "application/png"
	}else if request.Type=="jpg"{
		contentType = "application/jpg"
	}

	filePath:="./../images/"+request.Name+"."+request.Type
	objectName :=request.Name+"."+request.Type
	err = decode.Decode(request.Info,filePath)
	if err != nil {
		return err
	}

	info, err := minioClient.FPutObject(ctx, request.Bucket, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		log.Println(err)
		return err
	}
	log.Printf("Successfully uploaded %s of size %d\n", request.Name+"."+request.Type, info.Size)
	return err
}