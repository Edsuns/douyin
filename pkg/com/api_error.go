package com

type APIError struct {
	code int32
	msg  string
}

func NewAPIError(code int32, msg string) *APIError {
	err := APIError{code: code, msg: msg}
	return &err
}

func (err *APIError) Error() string {
	return err.msg
}
