package dao

import (
	"database/sql"
	"gorm.io/gorm"
)

type VideoFavorite struct {
	Model
	ProfileUserId int64 `gorm:"primaryKey;autoIncrement:false" json:"-"`
	VideoId       int64 `gorm:"primaryKey;autoIncrement:false" json:"-"`
	ProfileUser   User  `json:"-"`
	Video         Video `json:"-"`
}

func GetProfileVideos(userId int64) []*Video {
	var (
		videos   []*Video
		videoIds []int64
	)
	db.Select("video_id").Where("profile_user_id=?", userId).Find(&videoIds)
	err := db.Preload(
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

func addFavoriteVideo(tx *gorm.DB, userId, videoId int64) (err error) {
	var videoFavorite VideoFavorite

	tx.Unscoped().First(&videoFavorite,
		"profile_user_id = ? and video_id = ?",
		userId, videoId)
	// if there is an undeleted record, no need to add
	if videoFavorite.VideoId > 0 && !videoFavorite.DeletedAt.Valid {
		return nil
	}

	// increase FavoriteCount with optimistic lock
	err = addFavoriteCount(tx, videoId, 1)
	if err != nil {
		return err
	}

	// if there is a record and soft-deleted, set deleted false
	if videoFavorite.DeletedAt.Valid {
		videoFavorite.DeletedAt.Valid = false
		return tx.Unscoped().Updates(&videoFavorite).Error
	}

	// assign ids and update/insert
	videoFavorite.ProfileUserId = userId
	videoFavorite.VideoId = videoId
	return tx.Save(&videoFavorite).Error
}

func RemoveFavoriteVideo(userId, videoId int64) error {
	return db.Transaction(func(tx *gorm.DB) error {
		return removeFavoriteVideo(tx, userId, videoId)
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted})
}

func removeFavoriteVideo(tx *gorm.DB, userId, videoId int64) (err error) {
	var videoFavorite VideoFavorite
	tx.Unscoped().First(&videoFavorite,
		"profile_user_id = ? and video_id = ?",
		userId, videoId)
	// if no records or record is soft-deleted, no need to delete
	if videoFavorite.ProfileUserId <= 0 || videoFavorite.DeletedAt.Valid {
		return nil
	}

	// decrease FavoriteCount with optimistic lock
	err = addFavoriteCount(tx, userId, -1)
	if err != nil {
		return err
	}

	// soft delete the record
	return tx.Delete(&videoFavorite).Error
}
