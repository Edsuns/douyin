package api

import (
	"douyin/app/errs"
	"douyin/app/service"
	"douyin/pkg/com"
	"douyin/pkg/security"
	"douyin/pkg/validate"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type FavoriteRequest struct {
	VideoId    int64 `form:"video_id"`
	ActionType int   `form:"action_type" validate:"required,oneof=1 2"`
}

type FavoriteResponse struct {
}

// FavoriteAction adds a favorite video
func FavoriteAction(c *gin.Context) {
	myUserId := security.GetUserId(c)
	rq := validate.StructQuery(c, &FavoriteRequest{})
	if rq == nil {
		return
	}
	res := service.ChangeFavorite(myUserId, rq.VideoId, rq.ActionType)
	if res == true {
		com.SuccessStatus(c)
	} else {
		com.Error(c, errs.UserNotFound)
	}
}

// FavoriteList returns favorite video list
func FavoriteList(c *gin.Context) {
	myUserId := security.GetUserId(c)

	userId, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if err != nil {
		c.Status(http.StatusBadRequest)
	}

	favorites := service.GetFavorite(userId)
	for _, f := range favorites {
		f.IsFavorite = service.IsFavorite(f.ID, myUserId)
	}
	com.Success(c, &VideoListResponse{
		VideoList: favorites,
	})
}
