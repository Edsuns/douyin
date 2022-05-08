package dao

import (
	"douyin/app/config"
	"douyin/pkg/database"
	"gorm.io/gorm"
)

var db *gorm.DB

func Setup() {
	db = database.Init(config.Val.Gorm.LogLevel, config.Val.Mysql)

	err := db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User{})
	if err != nil {
		panic(err)
	}
}
