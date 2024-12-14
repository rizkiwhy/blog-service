package main

import (
	"fmt"
	"net/http"
	"os"
	"rizkiwhy-blog-service/api/router"
	"rizkiwhy-blog-service/util/database"
	"rizkiwhy-blog-service/util/logger"

	pkgUser "rizkiwhy-blog-service/package/user"
	mUser "rizkiwhy-blog-service/package/user/model"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func main() {
	g := gin.Default()

	logger.InitLogger()

	router.SetupPingRoutes(g)
	db, err := database.MySQLConnection()
	if err != nil {
		log.Fatal().Err(err).Msg("[main] Failed to connect to MySQL")
	}
	db.AutoMigrate(&mUser.User{})

	redisClient, err := database.RedisConnection()
	if err != nil {
		log.Fatal().Err(err).Msg("[main] Failed to connect to Redis")
	}

	userRepository := pkgUser.NewRepository(db)
	userCacheRepository := pkgUser.NewCacheRepository(redisClient)
	userService := pkgUser.NewService(userRepository, userCacheRepository)
	router.SetupUserRoutes(g, userService)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", os.Getenv("SERVICE_PORT")),
		Handler: g,
	}

	server.ListenAndServe()
}
