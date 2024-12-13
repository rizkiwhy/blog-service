package router

import (
	"rizkiwhy-blog-service/api/handler"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	pingHandler := handler.NewPingHandler()
	router.GET("/ping", pingHandler.Ping)

	// More routes can be added here
}
