package server

import (
	"github.com/gin-gonic/gin"
	"local-object-storage/pkg/server/controller"
	"net/http"
)

func Server(controller controller.Controller) *gin.Engine{
	Server := gin.Default()
	Server.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello world",
		})
	})

	//バケット一覧の確認
	Server.GET("/read/bucket", controller.BucketListHandler())
	//バケットの作成
	Server.POST("/create/bucket",controller.CreateBucketHandler())
	//画像のアップロード
	Server.POST("/upload/images",controller.UpLoadHandler())

	return Server
}
