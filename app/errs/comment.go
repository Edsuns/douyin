package errs

import "douyin/pkg/com"

const (
	CodeActionTypeNotFound = BaseComment + iota
	CodeCommentIdNotFound
	CodeCommentLengthIsTooLong
	CodeVideoIdNotFound
)

var (
	ActionTypeNotFound     = com.NewAPIError(CodeActionTypeNotFound, "ActionType not found")
	CommentIdNotFound      = com.NewAPIError(CodeCommentIdNotFound, "commentId not found")
	CommentLengthIsTooLong = com.NewAPIError(CodeCommentLengthIsTooLong, "comment length is too long")
	VideoIdNotFound        = com.NewAPIError(CodeVideoIdNotFound, "videoId not found")
)
