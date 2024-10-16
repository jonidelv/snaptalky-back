package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"snaptalky/models"
	"snaptalky/utils"
	"snaptalky/utils/types"
)

type responsePayload struct {
	Message string `json:"message"`
	Tone    string `json:"tone" binding:"oneof=flirting friendly formal"`
}

func SaveResponse(c *gin.Context) {
	// Retrieve the existing user from the context
	user, err := getUserFromContext(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.ApiResponse{
			Status:  "error",
			Message: "failed to get user from context",
		})
		return
	}

	var responsePayload responsePayload

	if err := c.ShouldBindJSON(&responsePayload); err != nil {
		c.JSON(http.StatusBadRequest, types.ApiResponse{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	response := models.Response{UserID: user.ID, Message: responsePayload.Message, Tone: responsePayload.Tone}
	if err := response.Add(); err != nil {
		utils.LogError(err, "Error saving response")
		c.JSON(http.StatusInternalServerError, types.ApiResponse{
			Status:  "error",
			Message: "error saving response",
		})
		return
	}

	c.JSON(http.StatusOK, types.ApiResponse{
		Data:    struct{}{}, // Empty object
		Status:  "success",
		Message: "response saved successfully",
	})
}
