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

func loadVideosExtra(videos *[]*dao.Video, myUserId int64) {
	for _, v := range *videos {
		author := &v.Author
		author.IsFollow = service.IsFollowed(author.UserID, myUserId)

		v.IsFavorite = service.IsFavorite(v.ID, myUserId)
	}
}

// Feed returns video feed
func Feed(c *gin.Context) {
	myUserId := security.GetUserId(c)

	latestTime, err := strconv.ParseInt(c.Query("latest_time"), 10, 64)
	if err != nil || latestTime <= 0 {
		latestTime = time.Now().UnixMilli()
	}

	videos := service.GetVideoFeed(latestTime)
	loadVideosExtra(&videos, myUserId)

	var nextTime int64 = 0
	if len(videos) > 0 {
		nextTime = videos[len(videos)-1].CreatedAt.Unix()
	}
	com.Success(c, &FeedResponse{
		VideoList: videos,
		NextTime:  nextTime,
	})
}
