package errs

import "douyin/pkg/com"

const (
	CodeInvalidPwdAndUsr = BaseUser + iota
	CodeUserAlreadyExists
	CodeUserNotFound
)

var (
	InvalidPwdAndUsr  = com.NewAPIError(CodeInvalidPwdAndUsr, "invalid username and password")
	UserAlreadyExists = com.NewAPIError(CodeUserAlreadyExists, "user already exist")
	UserNotFound      = com.NewAPIError(CodeUserNotFound, "user not found")
)
