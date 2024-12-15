package presenter

import "net/http"

const (
	CreatePostSuccessMessage        = "Post created successfully"
	CreatePostFailureMessage        = "Failed to create post"
	CreatePostUnauthorizedMessage   = "Unauthorized to create post"
	CreatePostInvalidRequestMessage = "Invalid create post request"
	GetPostSuccessMessage           = "Post retrieved successfully"
	GetPostFailureMessage           = "Failed to retrieve post"
	UpdatePostUnauthorizedMessage   = "Unauthorized to update post request"
	UpdatePostFailureMessage        = "Failed to update post request error"
	UpdatePostSuccessMessage        = "Post updated successfully"
	UpdatePostNotFoundMessage       = "record not found"
)

var CreatePostStatusCodeMap = map[string]int{
	CreatePostFailureMessage:        http.StatusInternalServerError,
	CreatePostInvalidRequestMessage: http.StatusBadRequest,
}

var GetPostStatusCodeMap = map[string]int{
	GetPostFailureMessage: http.StatusInternalServerError,
}

var UpdatePostStatusCodeMap = map[string]int{
	UpdatePostFailureMessage:      http.StatusInternalServerError,
	UpdatePostNotFoundMessage:     http.StatusNotFound,
	UpdatePostUnauthorizedMessage: http.StatusUnauthorized,
}
