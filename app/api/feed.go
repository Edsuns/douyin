package api

import (
	"douyin/app/dao"
	"douyin/app/service"
	"douyin/pkg/com"
	"douyin/pkg/security"
	"github.com/gin-gonic/gin"
	"time"
)

type FeedResponse struct {
	com.Response
	VideoList []dao.Video `json:"video_list,omitempty"`
	NextTime  int64       `json:"next_time,omitempty"`
}

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	myUserId := security.GetUserId(c)

	// TODO
	videos := *service.GetVideoFeed(1)
	for i := 0; i < len(videos); i++ {
		author := &videos[i].Author
		author.IsFollow = service.IsFollowed(author.UserID, myUserId)
	}
	com.Success(c, &FeedResponse{
		VideoList: videos,
		NextTime:  time.Now().Unix(),
	})
}
