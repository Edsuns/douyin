package dao

import (
	"database/sql"
	"fmt"
	"gorm.io/gorm"
)

type VideoFavorite struct {
	Model
	ProfileUserId int64 `gorm:"primaryKey;autoIncrement:false" json:"-"`
	VideoId       int64 `gorm:"primaryKey;autoIncrement:false" json:"-"`
	ProfileUser   User  `json:"-"`
	Video         Video `json:"-"`
}

func GetProfileVideos(userId int64) *[]Video {
	var (
		videos   *[]Video
		videoIds []int64
	)
	db.Debug().Select("video_id").Where("profile_user_id=?", userId).Find(&videoIds)
	err := db.Debug().Preload(
		"Author").Preload(
		"File").Preload(
		"Cover").Find(&videos, "videos.id in?", videoIds).Error
	if err != nil {
		return nil
	}
	return videos
}

func AddFavoriteVideo(userid, videoId int64) error {
	return db.Transaction(func(tx *gorm.DB) error {
		return addFavoriteVideo(tx, userid, videoId)
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted})

}
func addFavoriteVideo(tx *gorm.DB, userid, videoId int64) error {
	var videoFavorite VideoFavorite
	videoFavorite.ProfileUserId = userid
	videoFavorite.VideoId = videoId
	//1、在video——favorite中添加，
	if err := db.Unscoped().Model(&VideoFavorite{}).Update("deleted_at", nil).Error; err != nil {
		if err := tx.Create(&videoFavorite).Error; err != nil {
			return err
		}
	}
	//2、在video中使videofavorite_count+1

	//update douyin_test_db.videos set videos.favorite_count = favorite_count+1 where id= 5;
	if err := tx.Raw("update videos set favorite_count = favorite_count+1 where id=?", videoId).Error; err != nil {
		return err
	}
	return nil
}

func RemoveFavoriteVideo(userId, videoId int64) error {
	return db.Transaction(func(tx *gorm.DB) error {
		return removeFavoriteVideo(tx, userId, videoId)
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted})
}

func removeFavoriteVideo(tx *gorm.DB, userId, videoId int64) error {

	if err := tx.Where("profile_user_id =? and video_id=?", userId, videoId).Delete(&VideoFavorite{}).Error; err != nil {
		fmt.Println(err)
		return err
	}
	if err := tx.Raw("update videos set favorite_count = favorite_count-1 where id=?", videoId).Error; err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
