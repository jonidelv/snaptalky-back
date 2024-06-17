package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"snaptalky/routes"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, routes.ApiResponse{
				Status:  "error",
				Message: "Authorization header is required",
			})
			c.Abort()
			return
		}

		appToken := os.Getenv("APP_TOKEN")
		if token != appToken {
			c.JSON(http.StatusUnauthorized, routes.ApiResponse{
				Status:  "error",
				Message: "Invalid token",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
