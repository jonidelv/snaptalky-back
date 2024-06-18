package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"os"
	"snaptalky/database"
	"snaptalky/models"
	"snaptalky/utils"
	"snaptalky/utils/types"
)

type startRequest struct {
	DeviceID string `json:"device_id"`
}

func StartApp(c *gin.Context) {
	var req startRequest
	if err := c.BindJSON(&req); err != nil {
		utils.LogError(err, "Invalid request payload")
		c.JSON(http.StatusBadRequest, types.ApiResponse{
			Status:  "error",
			Message: "Invalid request",
		})
		return
	}

	tknHeader := c.GetHeader("tkn")
	startToken := os.Getenv("START_TOKEN")
	if tknHeader != startToken {
		c.JSON(http.StatusUnauthorized, types.ApiResponse{
			Status:  "error",
			Message: "Unauthorized",
		})
		return
	}

	deviceID := req.DeviceID
	if deviceID == "" {
		c.JSON(http.StatusBadRequest, types.ApiResponse{
			Status:  "error",
			Message: "device_id is required",
		})
		return
	}

	appToken := os.Getenv("APP_TOKEN")
	if appToken == "" {
		utils.LogError(nil, "APP_TOKEN not set in environment")
		c.JSON(http.StatusInternalServerError, types.ApiResponse{
			Status:  "error",
			Message: "APP_TOKEN not set in environment",
		})
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
				utils.LogError(err, "Error creating user")
				c.JSON(http.StatusInternalServerError, types.ApiResponse{
					Status:  "error",
					Message: "Error creating user",
				})
				return
			}
		} else {
			utils.LogError(err, "Database error")
			c.JSON(http.StatusInternalServerError, types.ApiResponse{
				Status:  "error",
				Message: "Database error",
			})
			return
		}
	}

	c.JSON(http.StatusOK, types.ApiResponse{
		Status:  "success",
		Message: "User retrieved/created successfully",
		Data: gin.H{
			"user": user,
			"id":   appToken, // This is the token but we call it id
		},
	})
}
