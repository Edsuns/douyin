package api

import (
	"douyin/app/dao"
	"douyin/app/errs"
	"douyin/pkg/com"
	"douyin/pkg/security"
	"douyin/pkg/validate"
	"github.com/gin-gonic/gin"
)

type FavoriteRequest struct {
	videoId    int64 `form:"video_id"`
	actionType int   `form:"action_type"`
}

// FavoriteAction no practical effect, just check if token is valid
func FavoriteAction(c *gin.Context) {
	myUserId := security.GetUserId(c)

	rq := validate.StructQuery(c, &FavoriteRequest{})
	if rq == nil {
		return
	}

	if myUserId > 0 {
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
