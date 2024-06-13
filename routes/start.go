package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"os"
	"snaptalky/database"
	"snaptalky/models"
)

type startRequest struct {
	DeviceID string `json:"device_id"`
}

func StartApp(c *gin.Context) {
	var req startRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	tknHeader := c.GetHeader("tkn")
	startToken := os.Getenv("START_TOKEN")
	if tknHeader != startToken {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	deviceID := req.DeviceID
	if deviceID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "device_id is required"})
		return
	}

	appToken := os.Getenv("APP_TOKEN")
	if appToken == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "t not set in environment"})
		return
	}

	// Perform the upsert operation
	var user models.User
	err := database.DB.Where("device_id = ?", deviceID).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// Create a new user if not found
			user = models.User{
				DeviceID: deviceID,
			}
			if err := database.DB.Create(&user).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating user"})
				return
			}
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"user":  user,
		"token": appToken,
	})
}
