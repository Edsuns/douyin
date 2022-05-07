package dao

import (
	"douyin/app/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"strings"
)

var db *gorm.DB

func Setup() {
	connect(getGormLogLevel())

	err := db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User{})
	if err != nil {
		panic(err)
	}
}

func getGormLogLevel() logger.LogLevel {
	var logLevel logger.LogLevel
	switch strings.ToLower(config.Val.Gorm.LogLevel) {
	case "silent":
		logLevel = logger.Silent
	case "error":
		logLevel = logger.Error
	case "warn":
		logLevel = logger.Warn
	case "info":
	default:
		logLevel = logger.Info
	}
	return logLevel
}

func connect(logLevel logger.LogLevel) {
	c := config.Val.Mysql
	url := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		c.Username, c.Password, c.Host, c.Database)
	var err error
	db, err = gorm.Open(mysql.Open(url), &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	})
	if err != nil {
		panic("database connection failed")
	}
}
