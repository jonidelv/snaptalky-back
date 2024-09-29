package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"snaptalky/utils"
	"snaptalky/utils/openai"
	"snaptalky/utils/types"
)

type DataRequest struct {
	Text        string `json:"text"`
	Image       string `json:"image"`
	Tone        string `json:"tone" binding:"oneof=flirting friendly formal"`
	Context     string `json:"context"`
	ContextText string `json:"customToneText"`
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
		Image:   data.Image,
		Context: data.Context,
		Tone:    data.Tone,
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

	if err := retrieveAndIncrementScanCount(c); err != nil {
		utils.LogError(err, "Failed to process user count")
	}

	c.JSON(http.StatusOK, types.ApiResponse{
		Status:  "success",
		Message: "Request processed successfully",
		Data:    response,
	})
}

// retrieveAndIncrementScanCount retrieves a user by ID and increments their scan count.
func retrieveAndIncrementScanCount(c *gin.Context) error {
	user, err := getUserFromContext(c)
	if err != nil {
		utils.LogError(err, "Error getting user from context for increment scan count")
		return err
	}

	if err := user.IncrementScanCount(); err != nil {
		utils.LogError(err, "Error incrementing user scan count")
		return err
	}

	return nil
}
