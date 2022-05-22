package dao

import (
	"gorm.io/gorm"
)

type Video struct {
	Model
	ID            int64      `gorm:"primary_key" json:"id,omitempty"`
	AuthorID      int64      `gorm:"index" json:"-"`
	Author        Profile    `json:"author"`
	Title         string     `json:"title"`
	FileID        int64      `json:"-"`
	File          MediaFile  `json:"-"`
	CoverID       int64      `json:"-"`
	Cover         MediaFile  `json:"-"`
	FavoriteCount int64      `json:"favorite_count,omitempty"`
	CommentCount  int64      `json:"comment_count,omitempty"`
	Favorites     []*Profile `gorm:"many2many:video_favorites;" json:"-"`
	Comments      []Comment  `gorm:"foreignKey:VideoID;references:ID" json:"-"`

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
		return tx.Create(&Video{
			Title:    title,
			AuthorID: authorId,
			File:     *video,
			Cover:    *cover,
		}).Error
	})
}

func GetVideosByCreatedAt(after int64) *[]Video {
	var videos []Video
	err := db.Preload(
		"Author").Preload(
		"File").Preload(
		"Cover").Find(&videos).Error
	if err != nil {
		return nil
	}
	return &videos
}
