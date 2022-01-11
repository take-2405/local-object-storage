package handler

import (
	"github.com/gin-gonic/gin"
	"local-object-storage/src/application"
	"local-object-storage/src/presentation"
	"log"
	"net/http"
)

type BucketHandler interface {
	BucketList() gin.HandlerFunc
	CreateBucket() gin.HandlerFunc
}

type bucketHandler struct {
	getBucketListsUseCase application.GetBucketListsUseCase
	createBucketUsecase   application.CreateBucketUseCase
}

// NewMinioHandler minioHandlerのコンストラクタ
func NewBucketHandler(getBucketListsUseCase application.GetBucketListsUseCase, createBucketUsecase application.CreateBucketUseCase) BucketHandler {
	return &bucketHandler{
		getBucketListsUseCase: getBucketListsUseCase,
		createBucketUsecase:   createBucketUsecase,
	}
}

func (bh *bucketHandler) BucketList() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, "ok")
	}
}

func (bh *bucketHandler) CreateBucket() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request presentation.CreateBucketRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			log.Println("[ERROR] Faild Bind JSON")
			c.JSON(http.StatusBadRequest, view.ReturnErrorResponse(
				http.StatusBadRequest,
				"Internal Server Error",
				"Request is error",
			))
			return
		}

		if err := dao.CreateBuckt(ctrl.Minio.Client, request.Name); err != nil {
			log.Println(err)
			if err.Error() == request.Name+"is already exist" {
				c.JSON(http.StatusBadRequest, view.ReturnErrorResponse(
					http.StatusBadRequest,
					"Bad request",
					"that name bucket is already created",
				))
			} else {
				c.JSON(http.StatusInternalServerError, view.ReturnErrorResponse(
					http.StatusInternalServerError,
					"Internal Server Error",
					"Failed to create bucket",
				))
			}
			return
		}
		c.JSON(http.StatusOK, "success create bucket")
	}
}
