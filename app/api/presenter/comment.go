package presenter

import "net/http"

const (
	CreateCommentSuccessMessage        = "Comment created successfully"
	CreateCommentFailureMessage        = "Failed to create comment"
	CreateCommentInvalidRequestMessage = "Invalid create comment request"
	CreateCommentUnauthorizedMessage   = "Unauthorized to create comment"
)

var CreateCommentStatusCodeMap = map[string]int{
	CreatePostFailureMessage:        http.StatusInternalServerError,
	CreatePostInvalidRequestMessage: http.StatusBadRequest,
}
