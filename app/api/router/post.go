package router

import (
	"rizkiwhy-blog-service/api/handler"
	"rizkiwhy-blog-service/api/middleware"
	pkgPost "rizkiwhy-blog-service/package/post"

	"github.com/gin-gonic/gin"
)

func SetupPostRoutes(r *gin.Engine, authMiddleware *middleware.AuthMiddleware, service pkgPost.Service) {
	postHandler := handler.NewPostHandler(service)
	postsRouter := r.Group("/posts")
	postsRouter.Use(authMiddleware.AuthJWT())
	{
		postsRouter.GET("/", postHandler.Search)
		postsRouter.GET("/:id", postHandler.GetByID)
		postsRouter.PUT("/:id", postHandler.Update)
		postsRouter.POST("/", postHandler.Create)
		postsRouter.DELETE("/:id", postHandler.Delete)
	}
}
