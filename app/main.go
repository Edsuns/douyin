package main

import (
	"douyin/app/config"
	"douyin/app/dao"
	"douyin/pkg/security"
	"douyin/pkg/validate"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"path/filepath"
)

func Setup(path string) *gin.Engine {
	config.Load(path)
	security.Setup(config.Val.Jwt)
	dao.Setup()
	validate.Setup()

	// make static filepath dirs
	staticDir := filepath.Join(path, config.Val.Static.Filepath)
	if _, err := os.Stat(staticDir); errors.Is(err, os.ErrNotExist) {
		// does not exist
		// mkdir
		err := os.MkdirAll(config.Val.Static.Filepath, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}

	return setupRouter()
}

func main() {
	err := Setup("./").Run(":" + config.Val.Port)
	if err != nil {
		log.Fatalln(err)
	}
}
