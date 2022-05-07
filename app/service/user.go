package service

import (
	"douyin/app/config"
	"douyin/app/dao"
	"douyin/app/errs"
	"douyin/pkg/util"
)

// Register create a new account
func Register(username, password string) (*dao.User, error) {
	// validate the params
	if err := util.NotEmpty(username, password); err != nil {
		return nil, err
	}
	// check if exists
	if dao.GetUserByUsername(username) != nil {
		return nil, errs.UserAlreadyExists
	}
	// insert into database if not exist
	saved, err := dao.SaveUser(username, password)
	if !saved {
		return nil, err
	}
	user := dao.GetUserByUsername(username)
	if user == nil {
		panic("impossible: can't get user after saved")
	}
	// new user successfully created
	return user, nil
}

// Login returns User and token if login successfully, or both nil
func Login(username, password string) (*dao.User, *string) {
	// check if user exists
	user := dao.GetUserByUsername(username)
	if user == nil {
		return nil, nil
	}
	// verify the password
	if util.VerifyPassword(password, user.Password) {
		// password is valid, returns token
		token := GetTokenForUser(user)
		return user, &token
	}
	// password is invalid, returns nil
	return nil, nil
}

// GetTokenForUser returns token signed for the User
func GetTokenForUser(user *dao.User) string {
	secret := []byte(config.Val.Jwt.Secret)
	token, err := util.GenerateJwt(user.ID,
		config.Val.Jwt.Issuer,
		config.Val.Jwt.ExpiresIn,
		secret)
	if err != nil {
		panic("failed to sign jwt")
	}
	return token
}

// GetUserInfo returns User
func GetUserInfo(userId int64) *dao.User {
	return dao.GetUserByUserId(userId)
}
