package dao

import (
	"douyin/pkg/dbx"
	"gorm.io/gorm"
	"gorm.io/plugin/optimisticlock"
)

type Video struct {
	Model
	ID            int64                  `gorm:"primary_key" json:"id,omitempty"`
	AuthorID      int64                  `gorm:"index" json:"-"`
	Author        Profile                `json:"author"`
	Title         string                 `json:"title"`
	FileID        int64                  `json:"-"`
	File          MediaFile              `json:"-"`
	CoverID       int64                  `json:"-"`
	Cover         MediaFile              `json:"-"`
	FavoriteCount *int64                 `json:"favorite_count,omitempty"`
	CommentCount  *int64                 `json:"comment_count,omitempty"`
	Favorites     []*Profile             `gorm:"many2many:video_favorites;" json:"-"`
	Comments      []Comment              `gorm:"foreignKey:VideoID;references:ID" json:"-"`
	Version       optimisticlock.Version `json:"-"`

	// post-load
	IsFavorite bool   `gorm:"-" json:"is_favorite"`
	PlayUrl    string `gorm:"-" json:"play_url,omitempty"`
	CoverUrl   string `gorm:"-" json:"cover_url,omitempty"`
}

type MediaFile struct {
	Model
	ID   int64  `gorm:"primary_key" json:"id,omitempty"`
	Key  string `gorm:"uniqueIndex;not null;size:255" json:"key"`
	MIME string `gorm:"not null;size:127" json:"mime"`
	SHA1 string `gorm:"not null;size:40" json:"sha1"`
}

func SaveVideoFile(authorId int64, title string, video *MediaFile, cover *MediaFile) error {
	return db.Transaction(func(tx *gorm.DB) error {
		v := Video{
			Title:    title,
			AuthorID: authorId,
			File:     *video,
			Cover:    *cover,
		}
		var zero int64 = 0
		v.FavoriteCount = &zero
		v.CommentCount = &zero
		return tx.Create(&v).Error
	})
}

func GetVideosByAuthor(userId int64) (videos []*Video) {
	err := db.Preload(
		"Author").Preload(
		"File").Preload(
		"Cover").Order(
		"created_at desc").Find(&videos,
		"author_id = ?", userId).Error
	if err != nil {
		// TODO: log err
		return []*Video{}
	}
	return videos
}

func GetVideosByCreatedAtBefore(time int64) (videos []*Video) {
	err := db.Preload(
		"Author").Preload(
		"File").Preload(
		"Cover").Order(
		"created_at desc").Find(&videos,
		"unix_timestamp(created_at) < ?", time).Error
	if err != nil {
		// TODO: log err
		return []*Video{}
	}
	return videos
}

// addFavoriteCount with optimistic lock
func addFavoriteCount(tx *gorm.DB, videoId int64, amount int64) error {
	return dbx.SpinOptimisticLock(tx, videoId, func(video *Video) {
		if video.FavoriteCount == nil {
			var one int64 = 1
			video.FavoriteCount = &one
			return
		}
		*video.FavoriteCount += amount
	})
}

// addCommentCount with optimistic lock
func addCommentCount(tx *gorm.DB, videoId int64, amount int64) error {
	return dbx.SpinOptimisticLock(tx, videoId, func(video *Video) {
		if video.CommentCount == nil {
			var one int64 = 1
			video.CommentCount = &one
			return
		}
		*video.CommentCount += amount
	})
}
