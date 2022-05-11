package security

import (
	"github.com/gin-gonic/gin"
)

func GetBearerToken(c *gin.Context) string {
	token := c.GetHeader("Authorization")
	if len(token) <= 7 {
		return ""
	}
	return token[7:]
}
