package main

import (
	"net/http"
	"rizkiwhy-blog-service/api/router"

	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.Default()

	router.SetupRoutes(g)

	s := &http.Server{
		Addr:    ":8080",
		Handler: g,
	}

	s.ListenAndServe()
}
