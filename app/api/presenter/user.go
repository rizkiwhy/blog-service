package presenter

import (
	"net/http"
	"rizkiwhy-blog-service/package/user/model"
)

const (
	RegisterSuccessMessage        = "User registered successfully"
	RegisterFailureMessage        = "Failed to register user"
	RegisterInvalidRequestMessage = "Invalid registration request"
)

var RegisterStatusCodeMap = map[string]int{
	model.ErrEmailAlreadyExists: http.StatusConflict,
	model.ErrNotFound:           http.StatusNotFound,
	model.ErrInternalError:      http.StatusInternalServerError,
}
