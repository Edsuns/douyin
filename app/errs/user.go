package errs

import "errors"

var (
	UserAlreadyExists = errors.New("user already exist")
)
