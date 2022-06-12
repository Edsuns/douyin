package service

import (
	"douyin/app/dao"
)

func ChangeFavorite(userId, videoId int64, actionType int) bool {
	if actionType == 1 {
		err := dao.AddFavoriteVideo(userId, videoId)
		if err == nil {
			return true
		}
		return false
	} else if actionType == 2 {
		err := dao.RemoveFavoriteVideo(userId, videoId)
		if err == nil {
			return true
		}
		return false
	}
	return false
}

func GetFavorite(userId int64) []*dao.Video {
	videos := dao.GetFavoriteVideos(userId)
	if videos == nil {
		return []*dao.Video{}
	}
	loadStaticUrls(&videos)
	return videos
}

func IsFavorite(videoId, userId int64) bool {
	yes, _ := dao.HasFavorite(videoId, userId)
	return yes
}
