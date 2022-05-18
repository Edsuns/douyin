package dbx

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"strings"
)

type MysqlConfig struct {
	Host     string
	Database string
	Username string
	Password string
}

func Init(logLevel string, config MysqlConfig) *gorm.DB {
	return connect(parseGormLogLevel(logLevel), config)
}

func parseGormLogLevel(level string) logger.LogLevel {
	var logLevel logger.LogLevel
	switch strings.ToLower(level) {
	case "silent":
		logLevel = logger.Silent
	case "error":
		logLevel = logger.Error
	case "warn":
		logLevel = logger.Warn
	case "info":
		logLevel = logger.Info
	default:
		logLevel = logger.Info
	}
	return logLevel
}

func connect(logLevel logger.LogLevel, c MysqlConfig) *gorm.DB {
	url := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		c.Username, c.Password, c.Host, c.Database)
	db, err := gorm.Open(mysql.Open(url), &gorm.Config{
		Logger:                 logger.Default.LogMode(logLevel),
		SkipDefaultTransaction: true,
	})
	if err != nil {
		panic(err)
	}
	return db
}
