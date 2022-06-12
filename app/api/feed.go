package api

import (
	"douyin/app/dao"
	"douyin/app/service"
	"douyin/pkg/com"
	"douyin/pkg/security"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

type FeedResponse struct {
	com.Response
	VideoList []*dao.Video `json:"video_list,omitempty"`
	NextTime  int64        `json:"next_time,omitempty"`
}

// Feed returns video feed
func Feed(c *gin.Context) {
	myUserId := security.GetUserId(c)

	latestTime, err := strconv.ParseInt(c.Query("latest_time"), 10, 64)
	if err != nil {
		latestTime = time.Now().UnixMilli()
	}

	videos := service.GetVideoFeed(latestTime)
	for i := 0; i < len(videos); i++ {
		author := &videos[i].Author
		author.IsFollow = service.IsFollowed(author.UserID, myUserId)
	}

	var nextTime int64 = 0
	if len(videos) > 0 {
		nextTime = videos[len(videos)-1].CreatedAt.Unix()
	}
	com.Success(c, &FeedResponse{
		VideoList: videos,
		NextTime:  nextTime,
	})
}
