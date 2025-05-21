package api

import (
	"github.com/cla-github/file_manager/api/middlewares"
	"github.com/cla-github/file_manager/internal/user/handler"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"error": nil,
			"data":  "Pong...",
		})
	})

	router.POST("/auth", user.Signin)
	router.POST("/auth/create-account", user.Signup)

	// testing protected middleware route
	router.GET("/me", middlewares.AuthMiddleware(), user.Me)

	return router
}
