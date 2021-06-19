package main

import (
	"local-object-storage/pkg/server"
	"local-object-storage/pkg/server/controller"
	"local-object-storage/pkg/server/model/dao"
	"log"
	"os"
)

func main() {
	//minioのコネクション作成
	minio := dao.New()

	controller:= controller.NewController(minio)

	if err:= server.Server(controller).Run(); err!=nil{
		log.Fatal(err)
		os.Exit(1)
	}
}