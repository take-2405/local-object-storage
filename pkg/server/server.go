package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"local-object-storage/pkg/server/controller"
	"net/http"
	"time"
)

func Server(controller controller.Controller) *gin.Engine {
	Server := gin.Default()
	Server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "accept", "origin", "Cache-Control", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		MaxAge: 15 * time.Second,
	}))

	Server.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello world",
		})
	})
	//バケット一覧の確認
	Server.GET("/read/bucket", controller.BucketListHandler())
	//バケットの作成
	Server.POST("/create/bucket", controller.CreateBucketHandler())
	//画像のアップロード
	Server.POST("/upload/images", controller.UpLoadHandler())

	return Server
}
