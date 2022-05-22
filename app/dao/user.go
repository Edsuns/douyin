package dao

import (
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

func SaveUserAndProfile(username, password string) (*User, error) {
	// start a transaction
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	var err error

	// insert User
	err = tx.Create(&User{
		Username: username,
		Password: security.EncodePassword(password),
	}).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	// query User ID
	var user User
	err = tx.First(&user, "username = ?", username).Error
	if user.ID == 0 || err != nil {
		tx.Rollback()
		return nil, errors.New("can't get user after inserted")
	}

	defaultImg := "https://douyin.com/favicon.ico"
	defaultSign := "nothing"

	// insert Profile
	err = tx.Create(&Profile{
		UserID:          user.ID,
		Name:            username,
		Avatar:          defaultImg,
		BackgroundImage: defaultImg,
		Signature:       defaultSign,
	}).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return &user, nil
}

func GetUserByUsername(username string) *User {
	var user User
	db.First(&user, "username = ?", username)
	if user.ID > 0 {
		return &user
	}
	return nil
}

func GetUserByUserId(userId int64) *User {
	var user User
	db.First(&user, userId)
	if user.ID > 0 {
		return &user
	}
	return nil
}
