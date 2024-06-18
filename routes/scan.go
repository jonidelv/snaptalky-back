package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"snaptalky/database"
	"snaptalky/models"
	"snaptalky/utils"
	"snaptalky/utils/openai"
	"snaptalky/utils/types"
)

type DataRequest struct {
	UserID  int    `json:"user_id"`
	Text    string `json:"text"`
	Context string `json:"context"`
	Tone    string `json:"tone"`
	Image   string `json:"image"`
}

func ProcessResponse(c *gin.Context) {
	var data DataRequest

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, types.ApiResponse{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	openAIData := types.DataToBuildResponses{
		Text:    data.Text,
		Context: data.Context,
		Tone:    data.Tone,
		Image:   data.Image,
	}

	response, err := openai.GenerateResponses(&openAIData)
	if err != nil {
		utils.LogError(err, "Error generating responses from OpenAI")
		c.JSON(http.StatusInternalServerError, types.ApiResponse{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	if err := retrieveAndIncrementScanCount(data.UserID); err != nil {
		utils.LogError(err, "Failed to process user count")
	}

	c.JSON(http.StatusOK, types.ApiResponse{
		Status:  "success",
		Message: "Request processed successfully",
		Data:    response,
	})
}

// retrieveAndIncrementScanCount retrieves a user by ID and increments their scan count.
func retrieveAndIncrementScanCount(userID int) error {
	var user models.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		utils.LogError(err, "Error retrieving user from database")
		return err
	}

	if err := user.IncrementScanCount(); err != nil {
		utils.LogError(err, "Error incrementing user scan count")
		return err
	}

	return nil
}
