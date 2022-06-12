package api

import (
	"douyin/app/dao"
	"douyin/app/errs"
	"douyin/app/service"
	"douyin/pkg/com"
	"douyin/pkg/security"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type VideoListResponse struct {
	com.Response
	VideoList []*dao.Video `json:"video_list"`
}

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	myUserId := security.GetUserId(c)

	title := c.PostForm("title")
	if title == "" {
		com.Error(c, errs.EmptyTitle)
		return
	}

	file, err := c.FormFile("data")
	if err != nil {
		com.Error(c, err)
		return
	}

	if err := service.PublishVideo(myUserId, title, file); err != nil {
		com.Error(c, err)
		return
	}

	com.SuccessStatus(c)
}

// PublishList returns publish video list by user
func PublishList(c *gin.Context) {
	myUserId := security.GetUserId(c)

	userId, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	videos := service.GetVideoPublishList(userId)
	for i := 0; i < len(videos); i++ {
		author := &videos[i].Author
		author.IsFollow = service.IsFollowed(author.UserID, myUserId)
	}
	com.Success(c, &VideoListResponse{
		VideoList: videos,
	})
}
