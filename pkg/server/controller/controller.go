package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"local-object-storage/pkg/server/model/dao"
	"local-object-storage/pkg/server/model/dto"
	"local-object-storage/pkg/server/view"
	"log"
	"net/http"
)

type Controller struct {
	Minio dao.Minio
}

func NewController(miniocontroller dao.Minio) Controller {
	return Controller{Minio: miniocontroller}
}

func (ctrl *Controller) BucketListHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		buckets, err := dao.BucketList()
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, view.ReturnErrorResponse(
				http.StatusInternalServerError,
				"Internal Server Error",
				"Failed to get buckets",
			))
			return
		}
		c.JSON(http.StatusOK, view.ReturnBucketListResponse(buckets))
	}
}

func (ctrl *Controller) CreateBucketHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request dto.CreateBucketRequest
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

func (ctrl *Controller) UpLoadHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		bucketName := c.Query("bucket")
		if bucketName == "" {
			log.Println("[ERROR] request bucket is err")
			c.JSON(http.StatusBadRequest, view.ReturnErrorResponse(
				http.StatusBadRequest,
				"Bad Request",
				"bucket is error",
			))
			return
		}

		form, _ := c.MultipartForm()
		files := form.File["a"]
		var imageNames []string
		imageName := ""

		buckets, err := dao.BucketList()
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, view.ReturnErrorResponse(
				http.StatusInternalServerError,
				"Internal Server Error",
				"Failed to get buckets",
			))
			return
		}
		for i, j := range buckets {
			if j == bucketName {
				break
			}
			if i == len(buckets)-1 {
				c.JSON(http.StatusInternalServerError, view.ReturnErrorResponse(
					http.StatusBadRequest,
					"Bad request",
					"Bucket is not exist.",
				))
				return
			}
		}

		for _, file := range files {
			u, err := uuid.NewRandom()
			if err != nil {
				fmt.Println(err)
			}
			uu := u.String()
			if err := dao.Upload(ctrl.Minio.Client, file, uu); err != nil {
				log.Println(err)
				c.JSON(http.StatusInternalServerError, view.ReturnErrorResponse(
					http.StatusInternalServerError,
					"Internal Server Error",
					"Failed to create bucket",
				))
				return
			}
			imageName = uu
			imageNames = append(imageNames, imageName)

		}
		c.JSON(http.StatusOK, "seccess")
	}
}
