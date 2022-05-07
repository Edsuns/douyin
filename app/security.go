package main

import (
	"douyin/app/api"
	"douyin/app/config"
	"douyin/app/errs"
	"douyin/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

var publicRoutes = make(map[string]bool)

// SecurityMiddleware filter unauthorized requests
func SecurityMiddleware(ctx *gin.Context) {
	// ignore public routes
	if publicRoutes[ctx.FullPath()] {
		return
	}
	// get bearer token
	token, err := util.GetBearerToken(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized,
			api.Response{StatusCode: 1, StatusMsg: err.Error()})
		return
	}
	// verify token and get user id
	userId, err := getUserIdFromToken(token)
	if err != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	// attach user id to context
	ctx.Set("userId", userId)
}

// PublicRoutes marks routes as public and ignore them when filtering
func PublicRoutes(routes ...string) {
	for _, val := range routes {
		publicRoutes[val] = true
	}
}

// getUserIdFromToken verifies token and returns user id
func getUserIdFromToken(token string) (int64, error) {
	secret := []byte(config.Val.Jwt.Secret)
	jwt, err := util.ParseJwt(token, secret)
	if err != nil {
		return 0, err
	}
	if util.IsJwtExpired(jwt) {
		return 0, errs.JwtExpired
	}
	return jwt.UserId, nil
}
