package errs

import (
	"douyin/pkg/com"
)

const (
	CodeInvalidPwdAndUsr = BaseUser + iota
	CodeUserAlreadyExists
	CodeUserNotFound

	CodeNotAllowedToFollowYourself
)

var (
	InvalidPwdAndUsr           = com.NewAPIError(CodeInvalidPwdAndUsr, "invalid username and password")
	UserAlreadyExists          = com.NewAPIError(CodeUserAlreadyExists, "user already exist")
	UserNotFound               = com.NewAPIError(CodeUserNotFound, "user not found")
	NotAllowedToFollowYourself = com.NewAPIError(CodeNotAllowedToFollowYourself, "not allowed to follow yourself")
)
