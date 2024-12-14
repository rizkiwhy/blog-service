package main

import (
	"net/http"
	"rizkiwhy-blog-service/api/router"
	"rizkiwhy-blog-service/util/database"
	"rizkiwhy-blog-service/util/logger"

	pkgUser "rizkiwhy-blog-service/package/user"
	mUser "rizkiwhy-blog-service/package/user/model"

	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.Default()

	logger.InitLogger()

	router.SetupPingRoutes(g)
	db, err := database.MySQLConnection()
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&mUser.User{})

	userRepository := pkgUser.NewRepository(db)
	userService := pkgUser.NewService(userRepository)
	router.SetupUserRoutes(g, userService)

	server := &http.Server{
		Addr:    ":8080",
		Handler: g,
	}

	server.ListenAndServe()
}
