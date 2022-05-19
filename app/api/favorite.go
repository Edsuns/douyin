package api

import (
	"douyin/app/dao"
	"douyin/app/errs"
	"douyin/pkg/com"
	"github.com/gin-gonic/gin"
)

// FavoriteAction no practical effect, just check if token is valid
func FavoriteAction(c *gin.Context) {
	token := c.Query("token")

	if _, exist := usersLoginInfo[token]; exist {
		com.SuccessStatus(c)
	} else {
		com.Error(c, errs.UserNotFound)
	}
}

// FavoriteList all users have same favorite video list
func FavoriteList(c *gin.Context) {
	com.Success(c, &VideoListResponse{
		VideoList: make([]dao.Video, 0),
	})
}
