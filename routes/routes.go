package routes

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	userRoutes := router.Group("/user")
	{
		userRoutes.GET("/:id", GetUser)
		userRoutes.PUT("/", UpdateUser)
	}

	scanRoutes := router.Group("/scan")
	{
		scanRoutes.POST("/", ProcessResponse)
	}
}
