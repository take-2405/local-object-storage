package main

import (
	"local-object-storage/pkg"
	"local-object-storage/pkg/model/dao"
	"local-object-storage/pkg/controller"
	"log"
	"os"
)

func main() {
	//minioのコネクション作成
	minio :=dao.New()

	controller:=controller.NewController(minio)

	if err:=pkg.Server(controller).Run(); err!=nil{
		log.Fatal(err)
		os.Exit(1)
	}
}