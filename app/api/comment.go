package api

import (
	"douyin/app/dao"
	"douyin/app/errs"
	"douyin/pkg/com"
	"douyin/pkg/security"
	"douyin/pkg/validate"
	"github.com/gin-gonic/gin"
)

type CommentRequest struct {
	videoId     int64  `form:"video_id"`
	actionType  int    `form:"action_type"`
	commentText string `form:"comment_text"`
	commentId   string `form:"comment_id"`
}

type CommentListResponse struct {
	com.Response
	CommentList []dao.Comment `json:"comment_list,omitempty"`
}

// CommentAction no practical effect, just check if token is valid
func CommentAction(c *gin.Context) {
	myUserId := security.GetUserId(c)

	rq := validate.StructQuery(c, &CommentRequest{})
	if rq == nil {
		return
	}

	if myUserId > 0 {
		com.SuccessStatus(c)
	} else {
		com.Error(c, errs.UserNotFound)
	}
}

// CommentList all videos have same demo comment list
func CommentList(c *gin.Context) {
	com.Success(c, &CommentListResponse{
		CommentList: make([]dao.Comment, 0),
	})
}
