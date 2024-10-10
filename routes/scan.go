package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"snaptalky/utils"
	"snaptalky/utils/openai"
	"snaptalky/utils/types"
)

func ProcessResponse(c *gin.Context) {
	var openAIData types.DataToBuildResponses

	if err := c.ShouldBindJSON(&openAIData); err != nil {
		c.JSON(http.StatusBadRequest, types.ApiResponse{
			Status:  "error",
			Message: err.Error(),
		})
		return
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
