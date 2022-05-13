package com

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	CodeUnknown = -1
	CodeSuccess = 0
	MsgSuccess  = "ok"
)

var emptyDataResponse = Response{StatusCode: CodeSuccess, StatusMsg: MsgSuccess}

type response interface {
	GetStatusCode() int32
	GetMsg() string
	setStatusCode(code int32)
	setMsg(msg string)
}

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

func (r Response) GetStatusCode() int32 {
	return r.StatusCode
}

func (r Response) GetMsg() string {
	return r.StatusMsg
}

func (r *Response) setStatusCode(code int32) {
	r.StatusCode = code
}

func (r *Response) setMsg(msg string) {
	r.StatusMsg = msg
}

func errToResponse(err error) *Response {
	var code int32
	var msg string

	var apiError *APIError
	if errors.As(err, &apiError) {
		code = apiError.code
		msg = apiError.msg
	} else {
		code = CodeUnknown
		msg = err.Error()
	}

	return &Response{StatusCode: code, StatusMsg: msg}
}

func SuccessStatus(c *gin.Context) {
	c.JSON(http.StatusOK, emptyDataResponse)
}

func Success(c *gin.Context, json response) {
	json.setStatusCode(CodeSuccess)
	if json.GetMsg() == "" {
		json.setMsg(MsgSuccess)
	}
	c.JSON(http.StatusOK, json)
}

func Error(c *gin.Context, err error) {
	c.JSON(http.StatusOK, errToResponse(err))
}
