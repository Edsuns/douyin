package dao

type Profile struct {
	UserID        int64      `gorm:"primary_key;autoIncrement:false" json:"id"`
	Name          string     `gorm:"size:63" json:"name"`
	FollowCount   int64      `json:"follow_count"`
	FollowerCount int64      `json:"follower_count"`
	Followers     []*Profile `gorm:"many2many:profile_followers;" json:"-"`
	Videos        []Video    `gorm:"foreignKey:AuthorID;references:UserID" json:"-"`
	Favorites     []*Video   `gorm:"many2many:video_favorites;" json:"-"`
	Comments      []Comment  `gorm:"foreignKey:AuthorID;references:UserID" json:"-"`
}

type ProfileFollower struct {
	ProfileUserID  int64
	FollowerUserID int64
}

func GetProfileByUserId(userId int64) *Profile {
	var profile Profile
	db.First(&profile, "user_id = ?", userId)
	if profile.UserID > 0 {
		return &profile
	}
	return nil
}

func HasFollower(userId, followerId int64) (bool, error) {
	var follower ProfileFollower
	err := db.First(&follower,
		"profile_user_id = ? and follower_user_id = ?",
		userId, followerId).Error
	if err != nil {
		return false, err
	}
	return follower.FollowerUserID > 0, nil
}
