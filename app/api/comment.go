package api

import (
	"douyin/app/dao"
	"douyin/app/errs"
	"douyin/app/service"
	"douyin/pkg/com"
	"douyin/pkg/security"
	"douyin/pkg/validate"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CommentRequest struct {
	VideoId     int64  `form:"video_id" validate:"required"`
	ActionType  int    `form:"action_type" validate:"required,oneof=1 2"`
	CommentText string `form:"comment_text"`
	CommentId   int64  `form:"comment_id"`
}

type CommentResponse struct {
	com.Response
	Comment *dao.Comment `json:"comment,omitempty"`
}

type CommentListResponse struct {
	com.Response
	CommentList []dao.Comment `json:"comment_list,omitempty"`
}

// CommentAction add or delete a comment
func CommentAction(c *gin.Context) {
	myUserId := security.GetUserId(c)

	rq := validate.StructQuery(c, &CommentRequest{})
	if rq == nil {
		return
	}

	if myUserId <= 0 {
		com.Error(c, errs.UserNotFound)
		return
	}

	comment, err := service.AddOrDeleteComment(myUserId, rq.VideoId, rq.CommentText, rq.ActionType, rq.CommentId)
	if err != nil {
		com.Error(c, err)
		return
	}
	if comment != nil {
		comment.Author.IsFollow = service.IsFollowed(comment.AuthorID, myUserId)
		comment.CreateDate = comment.CreatedAt.Format("01-02")
	}
	com.Success(c, &CommentResponse{
		Comment: comment,
	})
}

// CommentList get comment list by video id
func CommentList(c *gin.Context) {
	myUserId := security.GetUserId(c)

	videoId, err := strconv.ParseInt(c.Query("video_id"), 10, 64)
	if err != nil {
		com.Error(c, err)
		return
	}

	comments, err := service.GetComments(videoId)
	if err != nil {
		com.Error(c, err)
		return
	}

	for _, comment := range comments {
		comment.Author.IsFollow = service.IsFollowed(comment.AuthorID, myUserId)
		comment.CreateDate = comment.CreatedAt.Format("01-02")
	}

	com.Success(c, &CommentListResponse{
		CommentList: comments,
	})
}
