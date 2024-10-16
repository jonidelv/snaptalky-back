package routes

import (
	"errors"
	"github.com/gin-gonic/gin"
	"snaptalky/middlewares"
	"snaptalky/models"
	"snaptalky/utils/types"
)

func SetupRoutes(router *gin.Engine) {
	userRoutes := router.Group("/user")
	userRoutes.Use(middlewares.AuthMiddleware())
	userRoutes.GET("/:id", GetUser)
	userRoutes.PATCH("/:id", UpdateUser)

	scanRoutes := router.Group("/scan")
	scanRoutes.Use(middlewares.AuthMiddleware())
	scanRoutes.POST("", ProcessResponse)

	responseRoutes := router.Group("/response")
	responseRoutes.Use(middlewares.AuthMiddleware())
	responseRoutes.POST("", SaveResponse)

	startRoutes := router.Group("/start")
	startRoutes.POST("", StartApp)
}

func getUserFromContext(c *gin.Context) (models.User, error) {
	user, exists := c.Get("user")
	if !exists {
		return models.User{}, errors.New("user not found in context")
	}

	userTyped := user.(models.User)
	return userTyped, nil
}

func getAppUser(user models.User) types.AppUser {
	appUser := types.AppUser{
		ID:                 user.ID,
		DeviceID:           user.DeviceID,
		Age:                user.Age,
		Gender:             user.Gender,
		Bio:                user.Bio,
		PublicID:           user.PublicID,
		IsPremium:          user.IsPremium,
		CommunicationStyle: user.CommunicationStyle,
		ScanCount:          user.ScanCount,
	}

	return appUser
}
