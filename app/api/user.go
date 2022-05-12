package api

import (
	"douyin/app/service"
	"douyin/pkg/com"
	"douyin/pkg/security"
	"douyin/pkg/validate"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// usersLoginInfo use map to store user info, and key is username+password for demo
// user data will be cleared every time the server starts
// test data: username=zhanglei, password=douyin
var usersLoginInfo = map[string]User{
	"zhangleidouyin": {
		Id:            1,
		Name:          "zhanglei",
		FollowCount:   10,
		FollowerCount: 5,
		IsFollow:      true,
	},
}

type UserLoginResponse struct {
	com.Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token,omitempty"`
}

type UserResponse struct {
	com.Response
	User User `json:"user"`
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
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: com.Response{StatusCode: 1, StatusMsg: err.Error()},
		})
	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: com.Response{StatusCode: 0},
			UserId:   user.ID,
			Token:    service.GetTokenForUser(user),
		})
	}
}

func Login(c *gin.Context) {
	rq := validate.StructQuery(c, &UsrPwd{})
	if rq == nil {
		return
	}

	if user, token := service.Login(rq.Username, rq.Password); token != nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: com.Response{StatusCode: 0},
			UserId:   user.ID,
			Token:    *token,
		})
	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: com.Response{StatusCode: 1, StatusMsg: "invalid username and password"},
		})
	}
}

func UserInfo(c *gin.Context) {
	myUserId := security.GetUserId(c)
	userId, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if err != nil {
		userId = myUserId
	}

	if profile := service.GetUserInfo(userId); profile != nil {
		c.JSON(http.StatusOK, UserResponse{
			Response: com.Response{StatusCode: 0},
			User: User{
				Id:            profile.UserID,
				Name:          profile.Name,
				FollowCount:   profile.FollowCount,
				FollowerCount: profile.FollowerCount,
				IsFollow:      service.IsFollowed(userId, myUserId),
			},
		})
	} else {
		c.JSON(http.StatusOK, UserResponse{
			Response: com.Response{StatusCode: 1, StatusMsg: "user doesn't exist"},
		})
	}
}
