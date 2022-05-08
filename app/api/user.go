package api

import (
	"douyin/app/service"
	"douyin/pkg/security"
	"douyin/pkg/validate"
	"github.com/gin-gonic/gin"
	"net/http"
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
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Response
	User User `json:"user"`
}

type UsrPwd struct {
	Username string `validate:"required,min=4"`
	Password string `validate:"required,min=8"`
}

func Register(c *gin.Context) {
	rq := validate.Struct(c, &UsrPwd{})
	if rq == nil {
		return
	}

	user, err := service.Register(rq.Username, rq.Password)
	if err != nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: err.Error()},
		})
	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0},
			UserId:   user.ID,
			Token:    service.GetTokenForUser(user),
		})
	}
}

func Login(c *gin.Context) {
	rq := validate.Struct(c, &UsrPwd{})
	if rq == nil {
		return
	}

	if user, token := service.Login(rq.Username, rq.Password); token != nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0},
			UserId:   user.ID,
			Token:    *token,
		})
	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "invalid username and password"},
		})
	}
}

func UserInfo(c *gin.Context) {
	userId := security.GetUserId(c)

	if user := service.GetUserInfo(userId); user != nil {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 0},
			User: User{
				Id:   user.ID,
				Name: user.Username,
			},
		})
	} else {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}
