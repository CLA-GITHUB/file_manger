package middlewares

import (
	"os"
	"strings"

	"github.com/cla-github/file_manager/types"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// get the headers from the request and validate it
		header := c.GetHeader("Authorization")
		if header == "" {
			c.AbortWithStatusJSON(401, gin.H{
				"error": "Missing token.",
				"data":  nil,
			})

			return
		}

		// get the token from authroization "Bearer token"
		jwtToken := strings.Split(header, " ")
		if len(jwtToken) != 2 {
			c.AbortWithStatusJSON(401, gin.H{
				"error": "Invalid authorization headerj.",
				"data":  nil,
			})

			return
		}

		// get the jwt secret from env
		jwtSecret := []byte(os.Getenv("JWT_SECRET"))

		// validate token
		token, err := jwt.ParseWithClaims(jwtToken[1], &types.Claims{}, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})
		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(401, gin.H{
				"error": "Invalid or expired token.",
				"data":  nil,
			})
			return
		}

		claims, ok := token.Claims.(*types.Claims)
		if !ok {
			c.AbortWithStatusJSON(401, gin.H{
				"error": "Invalid claims.",
				"data":  nil,
			})
			return
		}

		c.Set("uid", claims.UserId)

		c.Next()
	}
}
