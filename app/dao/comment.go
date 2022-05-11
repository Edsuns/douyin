package dao

type Comment struct {
	Model
	Id       int64   `json:"id,omitempty"`
	VideoID  int64   `gorm:"index;not null;" json:"video_id"`
	AuthorID int64   `gorm:"index;not null;" json:"-"`
	Author   Profile `json:"user"`
	Content  string  `gorm:"not null;" json:"content,omitempty"`
}
