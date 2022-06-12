package service

import (
	"douyin/app/config"
	"douyin/app/dao"
	"fmt"
)

func GetVideoFeed(latestTime int64) []*dao.Video {
	videos := dao.GetVideosByCreatedAtBefore(latestTime)
	loadStaticUrls(&videos)
	return videos
}

func GetVideoPublishList(userId int64) []*dao.Video {
	videos := dao.GetVideosByAuthor(userId)
	loadStaticUrls(&videos)
	return videos
}

func loadStaticUrls(videos *[]*dao.Video) {
	v := *videos
	for i := 0; i < len(v); i++ {
		v[i].PlayUrl = toStaticUrl(v[i].File.Key)
		v[i].CoverUrl = toStaticUrl(v[i].Cover.Key)
	}
}

func toStaticUrl(fileKey string) string {
	static := config.Val.Static
	return fmt.Sprintf("%s%s/%s", static.BaseUrl, static.Route, fileKey)
}
