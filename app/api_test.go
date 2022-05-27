package main

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUnauthorized(t *testing.T) {
	router := Setup("../test/")

	// 未登录直接查询用户信息
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/douyin/user/", nil)
	router.ServeHTTP(w, req)

	// 服务端响应 401 未授权的
	assert.Equal(t, 401, w.Code)
}
