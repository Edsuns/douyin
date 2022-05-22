package errs

import "douyin/pkg/com"

const (
	CodeEmptyTitle = BasePublish + iota
)

var (
	EmptyTitle = com.NewAPIError(CodeEmptyTitle, "title is empty")
)
