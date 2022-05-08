package api

import (
	"douyin/pkg/com"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CommentListResponse struct {
	com.Response
	CommentList []Comment `json:"comment_list,omitempty"`
}

// CommentAction no practical effect, just check if token is valid
func CommentAction(c *gin.Context) {
	token := c.Query("token")

	if _, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, com.Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, com.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// CommentList all videos have same demo comment list
func CommentList(c *gin.Context) {
	c.JSON(http.StatusOK, CommentListResponse{
		Response:    com.Response{StatusCode: 0},
		CommentList: DemoComments,
	})
}
