package presenter

import "github.com/gin-gonic/gin"

func SuccessResponse(message string, data interface{}) gin.H {
	response := gin.H{
		"message": message,
		"status":  "success",
	}
	if data != nil {
		response["data"] = data
	}
	return response
}

func FailureResponse(title, message string) gin.H {
	return gin.H{
		"title":   title,
		"message": message,
		"status":  "error",
	}
}
