package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jonidelv/snaptalky-back/models"
	"github.com/jonidelv/snaptalky-back/utils"
	"github.com/jonidelv/snaptalky-back/utils/types"
	"net/http"
)

type responsePayload struct {
	Message string `json:"message"`
	Tone    string `json:"tone" binding:"oneof=flirting friendly formal"`
}

func SaveResponse(c *gin.Context) {
	// Retrieve the existing user from the context
	user, err := getUserFromContext(c)
	if err != nil {
		utils.LogError(err, "failed to get user from context", utils.Object{"path": "routes/response.go"})
		c.JSON(http.StatusInternalServerError, types.ApiResponse{
			Status:  "error",
			Message: "failed to get user from context",
		})
		return
	}

	var responsePayload responsePayload

	if err := c.ShouldBindJSON(&responsePayload); err != nil {
		utils.LogError(err, "failed to bind JSON payload", utils.Object{"path": "routes/response.go", "user": user.ID})
		c.JSON(http.StatusBadRequest, types.ApiResponse{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	response := models.Response{UserID: user.ID, Message: responsePayload.Message, Tone: responsePayload.Tone}
	if err := response.Add(); err != nil {
		utils.LogError(err, "error saving response", utils.Object{"path": "routes/response.go", "user": user.ID})
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
