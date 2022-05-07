package errs

import "errors"

var (
	UserAlreadyExists = errors.New("user already exist")
	JwtExpired        = errors.New("jwt expired")
)
