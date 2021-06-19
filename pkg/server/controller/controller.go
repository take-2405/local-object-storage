package controller

import (
	"github.com/gin-gonic/gin"
	"local-object-storage/pkg/server/model/dao"
	"local-object-storage/pkg/server/model/dto"
	"local-object-storage/pkg/server/view"
	"log"
	"net/http"
)

type Controller struct{
	Minio dao.Minio
}

func NewController(miniocontroller dao.Minio) Controller {
	return Controller{Minio: miniocontroller}
}

func (ctrl *Controller)BucketListHandler()gin.HandlerFunc{
	return func(c *gin.Context) {
		buckets,err := dao.BucketList()
		if err !=nil{
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

func (ctrl *Controller)CreateBucketHandler()gin.HandlerFunc{
	return func(c *gin.Context) {
		var request dto.CreateBucketRequest
		if err:=c.ShouldBindJSON(&request);err!=nil{
			log.Println("[ERROR] Faild Bind JSON")
			c.JSON(http.StatusBadRequest, view.ReturnErrorResponse(
				http.StatusBadRequest,
				"Internal Server Error",
				"Request is error",
			))
			return
		}
		if err := dao.CreateBuckt(ctrl.Minio.Client,request.Name);err!=nil{
			log.Println(err)
			if err.Error()==request.Name+"is already exist"{
				c.JSON(http.StatusBadRequest, view.ReturnErrorResponse(
					http.StatusBadRequest,
					"Bad request",
					"that name bucket is already created",
				))
			}else {
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

func (ctrl *Controller)UpLoadHandler()gin.HandlerFunc{
	return func(c *gin.Context) {
		var request dto.UploadImageRequest
		if err:=c.ShouldBindJSON(&request);err!=nil{
			log.Println("[ERROR] Faild Bind JSON")
			c.JSON(http.StatusBadRequest, view.ReturnErrorResponse(
				http.StatusBadRequest,
				"Internal Server Error",
				"Request is error",
			))
			return
		}

		if request.Type!="jpg" && request.Type!="png"{
			log.Println("[ERROR] request type is err")
			c.JSON(http.StatusBadRequest, view.ReturnErrorResponse(
				http.StatusBadRequest,
				"Bad Request",
				"Request's type is error",
			))
			return
		}

		buckets,err := dao.BucketList()
		if err !=nil{
			log.Println(err)
			c.JSON(http.StatusInternalServerError, view.ReturnErrorResponse(
				http.StatusInternalServerError,
				"Internal Server Error",
				"Failed to get buckets",
			))
			return
		}

		for i,j := range buckets{
			if j==request.Bucket{
				break
			}
			if i== len(buckets)-1{
				c.JSON(http.StatusInternalServerError, view.ReturnErrorResponse(
					http.StatusBadRequest,
					"Bad request",
					"Bucket is not exist.",
				))
				return
			}
		}

		if err := dao.Upload(ctrl.Minio.Client,request);err!=nil{
			log.Println(err)
			c.JSON(http.StatusInternalServerError, view.ReturnErrorResponse(
				http.StatusInternalServerError,
				"Internal Server Error",
				"Failed to create bucket",
			))
			return
		}
	c.JSON(http.StatusOK,"seccess")
	}
}