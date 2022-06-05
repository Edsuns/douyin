package dao

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
	pcomment := Comment{
		AuthorID: userId,
		VideoID:  videoId,
		Content:  content,
	}
	err = tx.Create(&pcomment).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	var comment Comment
	err = tx.First(&comment, "id = ?",
		pcomment.Id).First(&comment.Author,
		"author_id = ?", pcomment.AuthorID).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return &comment, nil
}

func DeleteComment(commentId int64) error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	err := tx.Where("id = ?", commentId).Delete(&Comment{}).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func GetComments(videoId int64) ([]Comment, error) {
	tx := db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	var comments []Comment
	err := db.Where("video_id = ?", videoId).Find(&comments).Error
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
