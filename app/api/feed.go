package api

import (
	"douyin/pkg/com"
	"github.com/gin-gonic/gin"
	"time"
)

type FeedResponse struct {
	com.Response
	VideoList []Video `json:"video_list,omitempty"`
	NextTime  int64   `json:"next_time,omitempty"`
}

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	com.Success(c, &FeedResponse{
		VideoList: DemoVideos,
		NextTime:  time.Now().Unix(),
	})
}
