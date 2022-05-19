package service

import (
	"douyin/app/config"
	"douyin/app/dao"
	"fmt"
)

func GetVideoFeed(nextTime int64) *[]dao.Video {
	// TODO
	videos := *dao.GetVideosByCreatedAt(nextTime)
	for i := 0; i < len(videos); i++ {
		videos[i].PlayUrl = toStaticUrl(videos[i].File.Key)
		videos[i].CoverUrl = toStaticUrl(videos[i].Cover.Key)
	}
	return &videos
}

func toStaticUrl(fileKey string) string {
	static := config.Val.Static
	return fmt.Sprintf("%s%s/%s", static.BaseUrl, static.Route, fileKey)
}
