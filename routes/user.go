package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"snaptalky/database"
	"snaptalky/models"
	"snaptalky/utils"
	"snaptalky/utils/types"
)

func GetUser(c *gin.Context) {
	//id := c.Param("id")
	appUser, err := getUserFromContext(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.ApiResponse{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, types.ApiResponse{
		Status:  "success",
		Message: "user retrieved successfully",
		Data:    appUser,
	})
}

func UpdateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.LogError(err, "error binding JSON to user model")
		c.JSON(http.StatusBadRequest, types.ApiResponse{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	if err := database.DB.Save(&user).Error; err != nil {
		utils.LogError(err, "Error saving user to database")
		c.JSON(http.StatusInternalServerError, types.ApiResponse{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	appUser := getAppUser(user)

	c.JSON(http.StatusOK, types.ApiResponse{
		Status:  "success",
		Message: "user updated successfully",
		Data:    appUser,
	})
}
