package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"local-object-storage/src/application"
	"log"
	"net/http"
)

type UploadHandler interface {
	UpLoadImage() gin.HandlerFunc
}

type uploadHandler struct {
	uploadImageUseCase application.UploadUseCase
}

// NewMinioHandler minioHandlerのコンストラクタ
func NewUploadHandler(uploadImageUseCase application.UploadUseCase) UploadHandler {
	return &uploadHandler{
		uploadImageUseCase: uploadImageUseCase,
	}
}

func (uh *uploadHandler) UpLoadImage() gin.HandlerFunc {
	return func(c *gin.Context) {
		bucketName := c.Query("bucket")
		if bucketName == "" {
			log.Println("[ERROR] request bucket is err")
			c.JSON(http.StatusBadRequest, ReturnErrorResponse(
				http.StatusBadRequest,
				"Bad Request",
				"bucket is error",
			))
			return
		}

		form, _ := c.MultipartForm()
		files := form.File["image"]
		var imageNames []string
		imageName := ""
		for _, file := range files {
			u, err := uuid.NewRandom()
			if err != nil {
				log.Println(err)
			}
			imageName = u.String()
			if err := uh.uploadImageUseCase.UploadImage(file, imageName, bucketName); err != nil {
				log.Println(err)
				c.JSON(http.StatusInternalServerError, ReturnErrorResponse(
					http.StatusInternalServerError,
					"Internal Server Error",
					"Failed to create bucket",
				))
				return
			}
			imageNames = append(imageNames, imageName)
		}
		c.JSON(http.StatusOK, imageName)
	}
}
