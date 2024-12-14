package router

import (
	"rizkiwhy-blog-service/api/handler"

	"github.com/gin-gonic/gin"
)

func SetupPingRoutes(r *gin.Engine) {

	pingHandler := handler.NewPingHandler()
	r.GET("/ping", pingHandler.Ping)
}
