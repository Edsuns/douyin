package dao

type Video struct {
	Model
	ID            int64      `gorm:"primary_key" json:"id,omitempty"`
	AuthorID      int64      `gorm:"index" json:"-"`
	Author        Profile    `json:"author"`
	FileID        int64      `json:"-"`
	File          MediaFile  `json:"-"`
	CoverID       int64      `json:"-"`
	Cover         MediaFile  `json:"-"`
	FavoriteCount int64      `json:"favorite_count,omitempty"`
	CommentCount  int64      `json:"comment_count,omitempty"`
	Favorites     []*Profile `gorm:"many2many:video_favorites;" json:"-"`
	Comments      []Comment  `gorm:"foreignKey:VideoID;references:ID" json:"-"`
}

type MediaFile struct {
	Model
	ID   int64  `gorm:"primary_key" json:"id,omitempty"`
	Key  string `gorm:"uniqueIndex;not null;size:255" json:"key"`
	MIME string `gorm:"not null;size:127" json:"mime"`
	SHA1 string `gorm:"not null;size:40" json:"sha1"`
}
