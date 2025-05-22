package api

import (
	"github.com/cla-github/file_manager/api/middlewares"
	"github.com/cla-github/file_manager/internal/user/handler"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/auth", user.Signin)
	router.POST("/auth/create-account", user.Signup)

	v1 := router.Group("/v1/api")
	{
		// test connection
		v1.GET("/ping", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"error": nil,
				"data":  "Pong...",
			})
		})

		// auth routes
		authRoutes := v1.Group("/auth")
		{
			authRoutes.POST("/", user.Signin)
			authRoutes.POST("/create-account", user.Signup)
		}

		// testing protected middleware route
		v1.GET("/me", middlewares.AuthMiddleware(), user.Me)
	}
	return router
}
