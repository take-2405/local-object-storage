package presentation

import (
	"github.com/gin-gonic/gin"
	"local-object-storage/src/presentation/handler"
	"net/http"
)

func InitRouting(s *gin.Engine, bh handler.BucketHandler, uh handler.UploadHandler) {
	s.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello world",
		})
	})

	//バケット一覧の確認
	s.GET("/bucket/read", bh.BucketList())
	//バケットの作成
	s.POST("/bucket/create", bh.CreateBucket())
	//画像のアップロード
	s.POST("/upload/images", uh.UpLoadImage())
}
