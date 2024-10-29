package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/jonidelv/snaptalky-back/constants"
	"github.com/jonidelv/snaptalky-back/database"
	"github.com/jonidelv/snaptalky-back/models"
	"github.com/jonidelv/snaptalky-back/utils"
	"github.com/jonidelv/snaptalky-back/utils/types"
	"gorm.io/gorm"
	"net/http"
	"os"
	"time"
)

type startRequest struct {
	DeviceID string `json:"deviceID"`
	Platform string `json:"platform"`
}

func StartApp(c *gin.Context) {
	var req startRequest
	if err := c.BindJSON(&req); err != nil {
		utils.LogError(err, "Invalid request payload")
		c.JSON(http.StatusBadRequest, types.ApiResponse{
			Status:  "error",
			Message: "invalid request",
		})
		return
	}

	tknHeader := c.GetHeader("ids") // This is the token coming from the app calling it ids ;)
	startToken := os.Getenv("START_TOKEN")
	if tknHeader != startToken {
		c.JSON(http.StatusUnauthorized, types.ApiResponse{
			Status:  "error",
			Message: "unauthorized",
		})
		return
	}

	deviceID := req.DeviceID
	if deviceID == "" {
		c.JSON(http.StatusBadRequest, types.ApiResponse{
			Status:  "error",
			Message: "deviceID is required",
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
				Platform: req.Platform,
			}
			if err := database.DB.Create(&user).Error; err != nil {
				utils.LogError(err, "Error creating user")
				c.JSON(http.StatusInternalServerError, types.ApiResponse{
					Status:  "error",
					Message: "error creating user",
				})
				return
			}
		} else {
			utils.LogError(err, "Database error")
			c.JSON(http.StatusInternalServerError, types.ApiResponse{
				Status:  "error",
				Message: "database error",
			})
			return
		}
	}

	appToken, err := createJWTToken(user.ID)
	if err != nil {
		utils.LogError(err, "APP_TOKEN creation failed")
		c.JSON(http.StatusInternalServerError, types.ApiResponse{
			Status:  "error",
			Message: "app token creation failed",
		})
		return
	}

	appUser := getAppUser(user)

	c.JSON(http.StatusOK, types.ApiResponse{
		Status:  "success",
		Message: "user retrieved/created successfully",
		Data: gin.H{
			"user": appUser,
			"ids":  appToken, // This is the token, but we call it ids ;)
		},
	})
}

func createJWTToken(userID uuid.UUID) (string, error) {
	tokenKey := os.Getenv("TOKEN_KEY")
	if tokenKey == "" {
		return "", fmt.Errorf("empty token key")
	}

	issuedAt := time.Now()
	expiresAt := issuedAt.Add(constants.TokenValidDays * 24 * time.Hour)
	claims := jwt.RegisteredClaims{
		Subject:   userID.String(),
		IssuedAt:  jwt.NewNumericDate(issuedAt),
		ExpiresAt: jwt.NewNumericDate(expiresAt),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(tokenKey))
}
