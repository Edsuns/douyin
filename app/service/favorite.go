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

func GetFavorite(userId int64) []dao.Video {
	videos := *dao.GetProfileVideos(userId)
	if videos == nil {
		return []dao.Video{}
	}
	for i := 0; i < len(videos); i++ {
		videos[i].PlayUrl = toStaticUrl(videos[i].File.Key)
		videos[i].CoverUrl = toStaticUrl(videos[i].Cover.Key)
	}
	return videos

}
