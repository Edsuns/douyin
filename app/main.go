package main

import (
	"douyin/app/config"
	"douyin/app/dao"
	"douyin/pkg/security"
	"douyin/pkg/validate"
	"log"
)

func init() {
	config.Setup("./config.yaml")
	security.Setup(config.Val.Jwt)
	dao.Setup()
	validate.Setup()
}

func main() {
	err := SetupRouter().Run(":" + config.Val.Port)
	if err != nil {
		log.Fatalln(err)
	}
}
