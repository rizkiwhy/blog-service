package router

import (
	"rizkiwhy-blog-service/api/handler"
	"rizkiwhy-blog-service/api/middleware"
	pkgPost "rizkiwhy-blog-service/package/post"

	"github.com/gin-gonic/gin"
)

func SetupPostRoutes(r *gin.Engine, authMiddleware *middleware.AuthMiddleware, service pkgPost.Service) {
	postHandler := handler.NewPostHandler(service)
	r.Use(authMiddleware.AuthJWT())
	r.POST("/post", postHandler.Create)
}
