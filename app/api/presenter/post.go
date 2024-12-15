package presenter

import "net/http"

const (
	CreatePostSuccessMessage        = "Post created successfully"
	CreatePostFailureMessage        = "Failed to create post"
	CreatePostUnauthorizedMessage   = "Unauthorized to create post"
	CreatePostInvalidRequestMessage = "Invalid create post request"
)

var CreatePostStatusCodeMap = map[string]int{
	CreatePostFailureMessage:        http.StatusInternalServerError,
	CreatePostSuccessMessage:        http.StatusCreated,
	CreatePostInvalidRequestMessage: http.StatusBadRequest,
}
