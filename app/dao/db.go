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
		&User{}, &Profile{}, &ProfileFollower{}, &Video{}, &MediaFile{}, &Comment{},
	)
	if err != nil {
		panic(err)
	}
}

func TruncateAllTables() {
	// truncate many to many first
	db.Exec("truncate table profile_followers")
	db.Exec("truncate table video_favorites")

	// truncate in specific order
	db.Exec("truncate table comments")
	db.Exec("truncate table videos")
	db.Exec("truncate table media_files")
	db.Exec("truncate table profiles")
	db.Exec("truncate table users")
}
