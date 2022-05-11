package dao

import (
	"douyin/app/config"
	"douyin/pkg/database"
	"gorm.io/gorm"
	"time"
)

type Model struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

var db *gorm.DB

func Setup() {
	db = database.Init(config.Val.Gorm.LogLevel, config.Val.Mysql)

	err := db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
		&User{}, &Profile{}, &Video{}, &MediaFile{}, &Comment{},
	)
	if err != nil {
		panic(err)
	}
}

func TruncateAllTables() {
	db.Exec("truncate table users")
}
