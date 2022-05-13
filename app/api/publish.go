package api

import (
	"douyin/app/service"
	"douyin/pkg/com"
	"douyin/pkg/security"
	"github.com/gin-gonic/gin"
)

type VideoListResponse struct {
	com.Response
	VideoList []Video `json:"video_list"`
}

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	myUserId := security.GetUserId(c)

	file, err := c.FormFile("data")
	if err != nil {
		com.Error(c, err)
		return
	}

	if err := service.PublishVideo(myUserId, file); err != nil {
		com.Error(c, err)
		return
	}

	com.SuccessStatus(c)
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	com.Success(c, &VideoListResponse{
		VideoList: DemoVideos,
	})
}
