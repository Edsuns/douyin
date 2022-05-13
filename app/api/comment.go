package api

import (
	"douyin/app/errs"
	"douyin/pkg/com"
	"github.com/gin-gonic/gin"
)

type CommentListResponse struct {
	com.Response
	CommentList []Comment `json:"comment_list,omitempty"`
}

// CommentAction no practical effect, just check if token is valid
func CommentAction(c *gin.Context) {
	token := c.Query("token")

	if _, exist := usersLoginInfo[token]; exist {
		com.SuccessStatus(c)
	} else {
		com.Error(c, errs.UserNotFound)
	}
}

// CommentList all videos have same demo comment list
func CommentList(c *gin.Context) {
	com.Success(c, &CommentListResponse{
		CommentList: DemoComments,
	})
}
