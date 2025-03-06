package middlewares

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"

	"github.com/jonidelv/snaptalky-back/database"
	"github.com/jonidelv/snaptalky-back/models"
	"github.com/jonidelv/snaptalky-back/utils"
	"github.com/jonidelv/snaptalky-back/utils/types"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, types.ApiResponse{
				Status:  "error",
				Message: "Authorization header is required",
			})
			c.Abort()
			return
		}

		getKey := func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("TOKEN_KEY")), nil
		}
		token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, getKey)

		if err != nil || !token.Valid {
			utils.LogError(err, "invalid token or errored when using ParseWithClaims", utils.Object{"path": "middlewares/middlewares.go"})
			c.JSON(http.StatusUnauthorized, types.ApiResponse{
				Status:  "error",
				Message: "Invalid token",
			})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(*jwt.RegisteredClaims)
		if !ok {
			utils.LogError(err, "invalid token claims or errored when using token.Claims", utils.Object{"path": "middlewares/middlewares.go"})
			c.JSON(http.StatusUnauthorized, types.ApiResponse{
				Status:  "error",
				Message: "Invalid token claims",
			})
			c.Abort()
			return
		}

		// Check for expiration
		if !claims.ExpiresAt.After(time.Now()) {
			c.JSON(http.StatusUnauthorized, types.ApiResponse{
				Status:  "error",
				Message: "Token has expired",
			})
			c.Abort()
			return
		}

		userID, err := uuid.Parse(claims.Subject)
		if err != nil {
			utils.LogError(err, "invalid userID when uuid.Parse", utils.Object{"path": "middlewares/middlewares.go"})
			c.JSON(http.StatusUnauthorized, types.ApiResponse{
				Status:  "error",
				Message: "Invalid user ID",
			})
			c.Abort()
			return
		}

		var user models.User
		if err := database.DB.First(&user, userID).Error; err != nil {
			utils.LogError(err, "error when getting user from db with userID", utils.Object{"path": "middlewares/middlewares.go"})
			c.JSON(http.StatusUnauthorized, types.ApiResponse{
				Status:  "error",
				Message: err.Error(),
			})
			c.Abort()
			return
		}

		// Set user in context
		c.Set("user", user)

		c.Next()
	}
}
