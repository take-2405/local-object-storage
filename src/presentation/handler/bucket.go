package handler

import (
	"github.com/gin-gonic/gin"
	"local-object-storage/src/application"
	"log"
	"net/http"
)

type BucketHandler interface {
	BucketList() gin.HandlerFunc
	CreateBucket() gin.HandlerFunc
}

type bucketHandler struct {
	bucketUseCase application.BucketUseCase
}

// NewMinioHandler minioHandlerのコンストラクタ
func NewBucketHandler(bucketUseCase application.BucketUseCase) BucketHandler {
	return &bucketHandler{
		bucketUseCase: bucketUseCase,
	}
}

func (bh *bucketHandler) CreateBucket() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request CreateBucketRequest
		err := c.ShouldBindJSON(&request)
		if err != nil || len(request.Name) < 3 {
			log.Println("[ERROR] Faild Bind JSON")
			c.JSON(http.StatusBadRequest, ReturnErrorResponse(
				http.StatusBadRequest,
				"Internal Server Error",
				"Request is error",
			))
			return
		}

		if err := bh.bucketUseCase.CreateBucket(request.Name); err != nil {
			log.Println(err)
			if err.Error() == request.Name+"is already exist" {
				c.JSON(http.StatusBadRequest, ReturnErrorResponse(
					http.StatusBadRequest,
					"Bad request",
					"that name bucket is already created",
				))
			} else {
				c.JSON(http.StatusInternalServerError, ReturnErrorResponse(
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

func (bh *bucketHandler) BucketList() gin.HandlerFunc {
	return func(c *gin.Context) {
		buckets, err := bh.bucketUseCase.GetBucketLists()
		if err != nil {
			log.Println("[ERROR] Faild Bind JSON")
			c.JSON(http.StatusBadRequest, ReturnErrorResponse(
				http.StatusBadRequest,
				"Internal Server Error",
				"Request is error",
			))
			return
		}
		c.JSON(http.StatusOK, buckets)
	}
}
