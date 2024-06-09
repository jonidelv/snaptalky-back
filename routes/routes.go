package routes

import (
	"github.com/gin-gonic/gin"
	"snaptalky/routes/scan"
	"snaptalky/routes/user"
)

func SetupRoutes(router *gin.Engine) {
  userRoutes := router.Group("/user")
  {
    userRoutes.GET("/:id", user.GetUser)
    userRoutes.PUT("/", user.UpdateUser)
  }

  scanRoutes := router.Group("/scan")
  {
    scanRoutes.POST("/", scan.ProcessResponse)
  }
}
