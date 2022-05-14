package dao

import (
	"douyin/app/config"
	"douyin/pkg/dbx"
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
	db = dbx.Init(config.Val.Gorm.LogLevel, config.Val.Mysql)

	err := db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
		&User{}, &Profile{}, &ProfileFollower{}, &Video{}, &MediaFile{}, &Comment{},
	)
	if err != nil {
		panic(err)
	}
}

func TruncateAllTables() {
	db.Exec("set foreign_key_checks = 0")
	defer db.Exec("set foreign_key_checks = 1")

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
