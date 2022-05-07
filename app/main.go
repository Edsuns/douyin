package main

import (
	"douyin/app/config"
	"douyin/app/dao"
	"log"
)

func main() {
	config.Setup("./config.yaml")
	dao.Setup()

	r := SetupRouter()

	err := r.Run(":" + config.Val.Port)
	if err != nil {
		log.Fatalln(err)
	}
}
