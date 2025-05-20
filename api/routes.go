package api

import (
	"github.com/cla-github/file_manager/api/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/ping", handlers.Ping)

	router.POST("/auth", handlers.Signin)
	router.POST("/auth/create-account", handlers.Signup)

	return router
}
