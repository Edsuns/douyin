package dao

import (
	"douyin/pkg/security"
)

type User struct {
	ID       int64  `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func SaveUser(username, password string) (bool, error) {
	result := db.Create(&User{
		Username: username,
		Password: security.EncodePassword(password),
	})
	return result.RowsAffected > 0, result.Error
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
	db.First(&user, "id = ?", userId)
	if user.ID > 0 {
		return &user
	}
	return nil
}
