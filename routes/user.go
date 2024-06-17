package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"snaptalky/database"
	"snaptalky/models"
	"snaptalky/utils"
	"strconv"
)

func GetUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.LogError(err, "Invalid ID format")
		c.JSON(http.StatusBadRequest, ApiResponse{
			Status:  "error",
			Message: "Invalid ID",
		})
		return
	}

	var user models.User
	if err := database.DB.First(&user, id).Error; err != nil {
		utils.LogError(err, "Error retrieving user from database")
		c.JSON(http.StatusInternalServerError, ApiResponse{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, ApiResponse{
		Status:  "success",
		Message: "User retrieved successfully",
		Data:    user,
	})
}

func UpdateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.LogError(err, "Error binding JSON to user model")
		c.JSON(http.StatusBadRequest, ApiResponse{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	if err := database.DB.Save(&user).Error; err != nil {
		utils.LogError(err, "Error saving user to database")
		c.JSON(http.StatusInternalServerError, ApiResponse{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, ApiResponse{
		Status:  "success",
		Message: "User updated successfully",
		Data:    user,
	})
}
