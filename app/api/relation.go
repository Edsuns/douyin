package api

import (
	"douyin/app/dao"
	"douyin/app/service"
	"douyin/pkg/com"
	"douyin/pkg/security"
	"douyin/pkg/validate"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserListResponse struct {
	com.Response
	UserList []*dao.Profile `json:"user_list"`
}

type FollowRequest struct {
	Token      string `form:"token"`
	ToUserId   int64  `form:"to_user_id" validate:"required,min=1"`
	ActionType int    `form:"action_type" validate:"required,oneof=1 2"`
}

// RelationAction no practical effect, just check if token is valid
func RelationAction(c *gin.Context) {
	rq := validate.StructQuery(c, &FollowRequest{})
	if rq == nil {
		return
	}
	userId := security.GetUserId(c)

	if err := service.Follow(rq.ToUserId, userId, rq.ActionType == 2); err != nil {
		com.Error(c, err)
	} else {
		com.SuccessStatus(c)
	}
}

// FollowList all users have same follow list
func FollowList(c *gin.Context) {
	myUserId := security.GetUserId(c)
	userId, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	if follows, err := dao.GetFollows(userId); err == nil {
		for _, f := range follows {
			f.IsFollow = userId == myUserId || service.IsFollowed(f.UserID, myUserId)
		}
		com.Success(c, &UserListResponse{
			UserList: follows,
		})
	} else {
		com.Error(c, err)
	}
}

// FollowerList all users have same follower list
func FollowerList(c *gin.Context) {
	myUserId := security.GetUserId(c)
	userId, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	if followers, err := dao.GetFollowers(userId); err == nil {
		for _, f := range followers {
			f.IsFollow = service.IsFollowed(f.UserID, myUserId)
		}
		com.Success(c, &UserListResponse{
			UserList: followers,
		})
	} else {
		com.Error(c, err)
	}
}
