package router

import (
	"rizkiwhy-blog-service/api/handler"
	"rizkiwhy-blog-service/api/middleware"
	pkgComment "rizkiwhy-blog-service/package/comment"
	pkgPost "rizkiwhy-blog-service/package/post"

	"github.com/gin-gonic/gin"
)

func SetupPostRoutes(r *gin.Engine, authMiddleware *middleware.AuthMiddleware, service pkgPost.Service, commentService pkgComment.Service) {
	postHandler := handler.NewPostHandler(service)
	postsRouter := r.Group("/posts")
	postsRouter.Use(authMiddleware.AuthJWT())
	{
		postsRouter.GET("/", postHandler.Search)
		postsRouter.POST("/", postHandler.Create)
		postIDRouter := postsRouter.Group("/:id")
		postIDRouter.GET("/:id", postHandler.GetByID)
		postIDRouter.PUT("/:id", postHandler.Update)
		postIDRouter.DELETE("/:id", postHandler.Delete)
		SetupCommentRoutes(postIDRouter, authMiddleware, commentService)
	}
}
