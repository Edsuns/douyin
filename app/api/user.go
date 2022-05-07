package api

import (
	"douyin/app/service"
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

func Register(c *gin.Context) {
	username, u := c.GetPostForm("username")
	password, p := c.GetPostForm("password")
	if !u || !p {
		c.Status(http.StatusBadRequest)
		return
	}

	user, err := service.Register(username, password)
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
	username, u := c.GetPostForm("username")
	password, p := c.GetPostForm("password")
	if !u || !p {
		c.Status(http.StatusBadRequest)
		return
	}

	if user, token := service.Login(username, password); token != nil {
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
	userId := c.GetInt64("userId")

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
