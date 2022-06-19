package dao

import (
	"douyin/pkg/dbx"
	"douyin/pkg/security"
	"errors"
)

type User struct {
	Model
	ID       int64   `gorm:"primary_key" json:"id"`
	Username string  `gorm:"uniqueIndex;not null;size:63" json:"username"`
	Password string  `gorm:"not null;size:60" json:"password"`
	Profile  Profile `json:"-"`
}

func SaveUserAndProfile(username, password string) (user *User, err error) {
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

	// insert User
	err = tx.Create(&User{
		Username: username,
		Password: security.EncodePassword(password),
	}).Error
	if err != nil {
		return nil, err
	}

	// query User ID
	err = tx.First(&user, "username = ?", username).Error
	if user.ID == 0 || err != nil {
		return nil, errors.New("can't get user after inserted")
	}

	defaultImg := "https://douyin.com/favicon.ico"
	defaultSign := "nothing"

	profile := Profile{
		UserID:          user.ID,
		Name:            username,
		Avatar:          defaultImg,
		BackgroundImage: defaultImg,
		Signature:       defaultSign,
	}
	var zero int64 = 0
	profile.FollowCount = &zero
	profile.FollowerCount = &zero
	// insert Profile
	err = tx.Create(&profile).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func GetUserByUsername(username string) (*User, error) {
	var user User
	err := db.First(&user, "username = ?", username).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func ExistsUserByUsername(username string) (bool, error) {
	return dbx.Exists(db, &User{}, "username = ?", username)
}

func GetUserByUserId(userId int64) *User {
	var user User
	db.First(&user, userId)
	if user.ID > 0 {
		return &user
	}
	return nil
}
