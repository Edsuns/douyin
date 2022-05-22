package api

import (
	"douyin/app/dao"
	"douyin/app/errs"
	"douyin/app/service"
	"douyin/pkg/com"
	"douyin/pkg/security"
	"douyin/pkg/validate"
	"github.com/gin-gonic/gin"
	"strconv"
)

type UserAuthResponse struct {
	com.Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token,omitempty"`
}

type UserResponse struct {
	com.Response
	User dao.Profile `json:"user"`
}

type UsrPwd struct {
	Username string `form:"username" validate:"required,min=4"`
	Password string `form:"password" validate:"required,min=8"`
}

func Register(c *gin.Context) {
	rq := validate.StructQuery(c, &UsrPwd{})
	if rq == nil {
		return
	}

	user, err := service.Register(rq.Username, rq.Password)
	if err != nil {
		com.Error(c, err)
	} else {
		com.Success(c, &UserAuthResponse{
			UserId: user.ID,
			Token:  service.GetTokenForUser(user),
		})
	}
}

func Login(c *gin.Context) {
	rq := validate.StructQuery(c, &UsrPwd{})
	if rq == nil {
		return
	}

	if user, token := service.Login(rq.Username, rq.Password); token != nil {
		com.Success(c, &UserAuthResponse{
			UserId: user.ID,
			Token:  *token,
		})
	} else {
		com.Error(c, errs.InvalidPwdAndUsr)
	}
}

func UserInfo(c *gin.Context) {
	myUserId := security.GetUserId(c)
	userId, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if err != nil {
		userId = myUserId
	}

	if profile := service.GetUserInfo(userId); profile != nil {
		profile.IsFollow = service.IsFollowed(userId, myUserId)
		com.Success(c, &UserResponse{
			User: *profile,
		})
	} else {
		com.Error(c, errs.UserNotFound)
	}
}
