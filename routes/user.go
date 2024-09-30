package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"snaptalky/database"
	"snaptalky/utils/types"
)

func GetUser(c *gin.Context) {
	//id := c.Param("id")
	user, err := getUserFromContext(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.ApiResponse{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	appUser := getAppUser(user)

	c.JSON(http.StatusOK, types.ApiResponse{
		Status:  "success",
		Message: "user retrieved successfully",
		Data:    appUser,
	})
}

type UpdateUserInput struct {
	Age                *int    `json:"age,omitempty"`
	Gender             *string `json:"gender,omitempty" binding:"oneof=male female other"`
	Bio                *string `json:"bio,omitempty"`
	CommunicationStyle *string `json:"communicationStyle,omitempty" binding:"oneof=default direct passive"`
}

func UpdateUser(c *gin.Context) {
	// Retrieve the existing user from the context
	user, err := getUserFromContext(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.ApiResponse{
			Status:  "error",
			Message: "failed to get user from context",
		})
		return
	}

	var input UpdateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, types.ApiResponse{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	// Apply updates to the user object
	if input.Age != nil {
		user.Age = *input.Age
	}
	if input.Gender != nil {
		user.Gender = *input.Gender
	}
	if input.Bio != nil {
		user.Bio = *input.Bio
	}
	if input.CommunicationStyle != nil {
		user.CommunicationStyle = *input.CommunicationStyle
	}

	if err := database.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, types.ApiResponse{
			Status:  "error",
			Message: "failed to update user in database",
		})
		return
	}

	appUser := getAppUser(user)

	// Return the updated user
	c.JSON(http.StatusOK, types.ApiResponse{
		Status:  "success",
		Message: "user updated successfully",
		Data:    appUser,
	})
}
