package handler

import (
	"github.com/gin-gonic/gin"
	"local-object-storage/src/application"
	"net/http"
)

type UploadHandler interface {
	UpLoadImage() gin.HandlerFunc
}

type uploadHandler struct {
	uploadImageUseCase application.UploadImageUseCase
}

// NewMinioHandler minioHandlerのコンストラクタ
func NewUploadHandler(uploadImageUseCase application.UploadImageUseCase) UploadHandler {
	return &uploadHandler{
		uploadImageUseCase: uploadImageUseCase,
	}
}

func (oh *uploadHandler) UpLoadImage() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, "ok")
	}
}

//func (oh *objectStorageHandler) BucketList() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		c.JSON(http.StatusOK, "ok")
//	}
//}
//
//func (oh *objectStorageHandler) CreateBucket() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		c.JSON(http.StatusOK, "ok")
//	}
//}
