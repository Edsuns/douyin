package dao

type Profile struct {
	UserID        int64      `gorm:"primary_key" json:"id"`
	Name          string     `gorm:"size:32" json:"name"`
	FollowCount   int64      `json:"follow_count"`
	FollowerCount int64      `json:"follower_count"`
	Followers     []*Profile `gorm:"many2many:profile_followers;" json:"-"`
	Videos        []Video    `gorm:"foreignKey:AuthorID;references:UserID" json:"-"`
	Favorites     []*Video   `gorm:"many2many:video_favorites;" json:"-"`
	Comments      []Comment  `gorm:"foreignKey:AuthorID;references:UserID" json:"-"`
}
