package router

import (
	"rizkiwhy-blog-service/api/handler"
	"rizkiwhy-blog-service/api/middleware"
	pkgComment "rizkiwhy-blog-service/package/comment"

	"github.com/gin-gonic/gin"
)

func SetupCommentRoutes(r *gin.RouterGroup, authMiddleware *middleware.AuthMiddleware, service pkgComment.Service) {
	commentHandler := handler.NewCommentHandler(service)
	commentsRouter := r.Group("/comments")
	commentsRouter.Use(authMiddleware.AuthJWT())
	{
		commentsRouter.POST("/", commentHandler.Create)
		commentsRouter.GET("/", commentHandler.Search)
	}
}
