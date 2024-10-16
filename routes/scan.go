package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"snaptalky/utils"
	"snaptalky/utils/openai"
	"snaptalky/utils/types"
)

func ProcessResponse(c *gin.Context) {
	user, err := getUserFromContext(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.ApiResponse{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	if !user.IsPremium && user.ScanCount > 2 {
		c.JSON(http.StatusForbidden, types.ApiResponse{
			Status:  "error",
			Message: "You are not premium",
		})
	}

	var openAIData types.DataToBuildResponses

	if err := c.ShouldBindJSON(&openAIData); err != nil {
		c.JSON(http.StatusBadRequest, types.ApiResponse{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	//var usages int
	response, usages, err := openai.GenerateResponses(&openAIData)
	if err != nil {
		utils.LogError(err, "Error generating responses from OpenAI")
		c.JSON(http.StatusInternalServerError, types.ApiResponse{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	if err := retrieveAndIncrementCounts(c, usages); err != nil {
		utils.LogError(err, "Failed to process user counts")
	}

	c.JSON(http.StatusOK, types.ApiResponse{
		Status:  "success",
		Message: "Request processed successfully",
		Data:    response,
	})
}

// retrieveAndIncrementCounts retrieves a user by ID and increments their counts.
func retrieveAndIncrementCounts(c *gin.Context, usages int) error {
	user, err := getUserFromContext(c)
	if err != nil {
		utils.LogError(err, "Error getting user from context for increment usages")
		return err
	}

	if err := user.IncrementCountsAndUsages(usages); err != nil {
		utils.LogError(err, "Error incrementing user usages count")
		return err
	}

	return nil
}
