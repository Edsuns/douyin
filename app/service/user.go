package service

import (
	"douyin/app/dao"
	"douyin/app/errs"
	"douyin/pkg/assert"
	"douyin/pkg/security"
)

// Register create a new account
func Register(username, password string) (user *dao.User, err error) {
	// validate the params
	if err = assert.NotEmpty(username, password); err != nil {
		return nil, err
	}
	// check if exists
	exists, err := dao.ExistsUserByUsername(username)
	if exists {
		return nil, errs.UserAlreadyExists
	}
	if err != nil {
		return nil, err
	}
	// insert into database if not exist
	user, err = dao.SaveUserAndProfile(username, password)
	return user, err
}

// Login returns User and token if login successfully, or both nil
func Login(username, password string) (*dao.User, *string) {
	// check if user exists
	user, err := dao.GetUserByUsername(username)
	if user == nil || err != nil {
		if err != nil {
			// TODO: log error
		}
		return nil, nil
	}
	// verify the password
	if security.VerifyPassword(password, user.Password) {
		// password is valid, returns token
		token := GetTokenForUser(user)
		return user, &token
	}
	// password is invalid, returns nil
	return nil, nil
}

// GetTokenForUser returns token signed for the User
func GetTokenForUser(user *dao.User) string {
	token, err := security.GenerateJwt(user.ID)
	if err != nil {
		panic("failed to sign jwt")
	}
	return token
}

// GetUserInfo returns User
func GetUserInfo(userId int64) *dao.Profile {
	return dao.GetProfileByUserId(userId)
}

func IsFollowed(userId, followerId int64) bool {
	if userId == followerId {
		return false
	}
	yes, _ := dao.HasFollower(userId, followerId)
	return yes
}

func Follow(userId, followerId int64, discard bool) error {
	if userId == followerId {
		return errs.NotAllowedToFollowYourself
	}
	var err error
	if discard {
		err = dao.RemoveFollower(userId, followerId)
	} else {
		err = dao.AddFollower(userId, followerId)
	}
	return err
}
