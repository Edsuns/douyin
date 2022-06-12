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

func SaveComment(userId int64, videoId int64, content string) (*Comment, error) {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	var err error

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
		tx.Rollback()
		return nil, err
	}

	var comment Comment
	err = tx.First(&comment, "id = ?",
		c.Id).First(&comment.Author,
		"author_id = ?", c.AuthorID).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return &comment, nil
}

func DeleteComment(commentId int64) error {
	return db.Transaction(func(tx *gorm.DB) (err error) {
		var comment Comment

		err = tx.Find(&comment, "id = ?", commentId).Error
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

func GetComments(videoId int64) (comments []*Comment, err error) {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	err = db.Where("video_id = ?", videoId).Find(&comments).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	for _, comment := range comments {
		err = db.Where("user_id = ?", comment.AuthorID).First(&comment.Author).Error
		if err != nil {
			return nil, err
		}
	}

	return comments, nil
}
