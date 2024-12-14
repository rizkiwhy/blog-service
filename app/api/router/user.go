package router

import (
	"rizkiwhy-blog-service/api/handler"
	pkgUser "rizkiwhy-blog-service/package/user"

	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(r *gin.Engine, service pkgUser.Service) {
	userHandler := handler.NewUserHandler(service)
	r.POST("/register", userHandler.Register)
	r.POST("/login", userHandler.Login)
}
