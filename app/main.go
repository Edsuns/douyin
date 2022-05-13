package main

import (
	"douyin/app/config"
	"douyin/app/dao"
	"douyin/pkg/security"
	"douyin/pkg/validate"
	"errors"
	"log"
	"os"
)

func init() {
	config.Setup("./config.yaml")
	security.Setup(config.Val.Jwt)
	dao.Setup()
	validate.Setup()

	// make static filepath dirs
	if _, err := os.Stat(config.Val.Static.Filepath); errors.Is(err, os.ErrNotExist) {
		// does not exist
		// mkdir
		err := os.MkdirAll(config.Val.Static.Filepath, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
}

func main() {
	err := SetupRouter().Run(":" + config.Val.Port)
	if err != nil {
		log.Fatalln(err)
	}
}
