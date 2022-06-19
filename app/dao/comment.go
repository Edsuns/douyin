package dao

import (
	"database/sql"
	"gorm.io/gorm"
)

type Comment struct {
	Model    `json:"-"`
	Id       int64   `json:"id,omitempty"`
	VideoID  int64   `gorm:"index;not null;" json:"video_id"`
	AuthorID int64   `gorm:"index;not null;" json:"-"`
	Author   Profile `json:"user"`
	Content  string  `gorm:"not null;" json:"content,omitempty"`

	// post-loads
	CreateDate string `gorm:"-" json:"create_date"`
}

func SaveComment(userId int64, videoId int64, content string) (comment *Comment, err error) {
	// start a transaction
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil || err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	if err = tx.Error; err != nil {
		return nil, err
	}

	// increase CommentCount with optimistic lock
	err = addCommentCount(tx, videoId, 1)
	if err != nil {
		return nil, err
	}

	c := Comment{
		AuthorID: userId,
		VideoID:  videoId,
		Content:  content,
	}
	err = tx.Create(&c).Error
	if err != nil {
		return nil, err
	}

	err = tx.First(&comment, "id = ?",
		c.Id).First(&comment.Author,
		"author_id = ?", c.AuthorID).Error
	if err != nil {
		return nil, err
	}

	return comment, nil
}

func DeleteComment(commentId int64, authorId int64) error {
	return db.Transaction(func(tx *gorm.DB) (err error) {
		var comment Comment

		err = tx.Find(&comment, "id = ? and author_id = ?", commentId, authorId).Error
		if err != nil {
			return err
		}
		// decrease CommentCount with optimistic lock
		err = addCommentCount(tx, comment.VideoID, -1)
		if err != nil {
			return err
		}

		err = tx.Delete(&comment).Error
		if err != nil {
			return err
		}

		return nil
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted})
}

func GetComments(videoId int64) ([]*Comment, error) {
	var comments []*Comment
	err := db.Where("video_id = ?", videoId).Preload(
		"Author").Find(&comments).Error
	if err != nil {
		return nil, err
	}
	return comments, nil
}
