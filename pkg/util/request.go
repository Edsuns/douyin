package util

import (
	"errors"
	"github.com/gin-gonic/gin"
)

var BearerTokenNotFound = errors.New("bearer token not found")

func GetBearerToken(c *gin.Context) (string, error) {
	token := c.GetHeader("Authorization")
	if len(token) <= 7 {
		return "", BearerTokenNotFound
	}
	return token[7:], nil
}
