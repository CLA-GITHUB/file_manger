package api

import (
	"github.com/cla-github/file_manager/api/handlers"
	"github.com/cla-github/file_manager/api/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/ping", handlers.Ping)

	router.POST("/auth", handlers.Signin)
	router.POST("/auth/create-account", handlers.Signup)

	// testing protected middleware route
	router.GET("/me", middlewares.AuthMiddleware(), handlers.Me)
	return router
}
