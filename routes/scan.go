package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jonidelv/snaptalky-back/models"
	"github.com/jonidelv/snaptalky-back/utils"
	"github.com/jonidelv/snaptalky-back/utils/openai"
	"github.com/jonidelv/snaptalky-back/utils/types"
	"net/http"
)

type scanRequest struct {
	Text              *string `json:"text"`
	Image             *string `json:"image"`
	Tone              string  `json:"tone" binding:"oneof=flirting friendly formal"`
	ResponseType      *string `json:"responseType"`
	AdditionalContext *string `json:"additionalContext"`
	Location          *string `json:"location"`
	Lang              *string `json:"lang"`
}

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

	var requestData scanRequest

	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, types.ApiResponse{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	// We need either text or image
	if (requestData.Text == nil || *requestData.Text == "") && (requestData.Image == nil || *requestData.Image == "") {
		c.JSON(http.StatusBadRequest, types.ApiResponse{
			Status:  "error",
			Message: "text or image is required",
		})
		return
	}

	if requestData.AdditionalContext != nil && len(*requestData.AdditionalContext) > 400 {
		c.JSON(http.StatusBadRequest, types.ApiResponse{
			Status:  "error",
			Message: "context is too large",
		})
		return
	}

	// We need either text or image
	if requestData.Tone == "" {
		c.JSON(http.StatusBadRequest, types.ApiResponse{
			Status:  "error",
			Message: "tone is needed it",
		})
		return
	}

	var previousResponses string
	previousResponses, prevRespErr := models.GetMessagesByTone(user.ID, requestData.Tone)
	if prevRespErr != nil {
		previousResponses = ""
	}

	Language := "English"
	if requestData.Lang != nil && *requestData.Lang != "" {
		Language = *requestData.Lang
	}

	openAIData := types.DataToBuildResponses{
		Text:              requestData.Text,
		Image:             requestData.Image,
		Tone:              requestData.Tone,
		AdditionalContext: requestData.AdditionalContext,
		Location:          requestData.Location,
		ResponseType:      requestData.ResponseType,
		Language:          Language,
		UserBio:           &user.Bio,
		UserGender:        &user.Gender,
		UserAge:           &user.Age,
		PreviousResponses: &previousResponses,
	}

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
