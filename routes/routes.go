package routes

import (
	"github.com/gin-gonic/gin"
	"snaptalky/middlewares"
)

func SetupRoutes(router *gin.Engine) {
	userRoutes := router.Group("/user")
	userRoutes.Use(middlewares.AuthMiddleware())
	userRoutes.GET("/:id", GetUser)
	userRoutes.PUT("/", UpdateUser)

	scanRoutes := router.Group("/scan")
	scanRoutes.Use(middlewares.AuthMiddleware())
	scanRoutes.POST("/", ProcessResponse)

	startRoutes := router.Group("/start")
	startRoutes.POST("/", StartApp)
}
