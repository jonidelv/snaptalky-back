package routes

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jonidelv/snaptalky-back/database"
	"github.com/jonidelv/snaptalky-back/utils"
	"github.com/jonidelv/snaptalky-back/utils/types"
	"net/http"
)

func GetUser(c *gin.Context) {
	//id := c.Param("id")
	user, err := getUserFromContext(c)
	if err != nil {
		utils.LogError(err, "failed to get user from context", utils.Object{"path": "routes/user.go"})
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
	Gender             *string `json:"gender,omitempty"`
	Bio                *string `json:"bio,omitempty"`
	CommunicationStyle *string `json:"communicationStyle,omitempty"`
	Lang               *string `json:"lang,omitempty"`
}

func UpdateUser(c *gin.Context) {
	// Retrieve the existing user from the context
	user, err := getUserFromContext(c)
	if err != nil {
		utils.LogError(err, "failed to get user from context", utils.Object{"path": "routes/user.go"})
		c.JSON(http.StatusInternalServerError, types.ApiResponse{
			Status:  "error",
			Message: "failed to get user from context",
		})
		return
	}

	var input UpdateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.LogError(err, "failed to bind JSON payload", utils.Object{"path": "routes/user.go", "user": user.ID})
		c.JSON(http.StatusBadRequest, types.ApiResponse{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	if input.Bio != nil && len(*input.Bio) > 350 {
		utils.LogError(errors.New("bio cannot be too large"), "bio > 350", utils.Object{"path": "routes/user.go", "user": user.ID})
		c.JSON(http.StatusBadRequest, types.ApiResponse{
			Status:  "error",
			Message: "bio cannot be too large",
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
	if input.Lang != nil {
		user.Lang = *input.Lang
	}
	if input.CommunicationStyle != nil {
		user.CommunicationStyle = *input.CommunicationStyle
	}

	if err := database.DB.Save(&user).Error; err != nil {
		utils.LogError(err, "failed to update user in database", utils.Object{"path": "routes/user.go", "user": user.ID})
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
