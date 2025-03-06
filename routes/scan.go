package routes

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/jonidelv/snaptalky-back/models"
	"github.com/jonidelv/snaptalky-back/utils"
	"github.com/jonidelv/snaptalky-back/utils/openai"
	"github.com/jonidelv/snaptalky-back/utils/types"
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
		utils.LogError(err, "failed to get user from context", utils.Object{"path": "routes/scan.go"})
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
		utils.LogError(err, "failed to bind JSON payload", utils.Object{"path": "routes/scan.go", "user": user.ID})
		c.JSON(http.StatusBadRequest, types.ApiResponse{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	// We need either text or image
	if (requestData.Text == nil || *requestData.Text == "") && (requestData.Image == nil || *requestData.Image == "") {
		utils.LogError(errors.New("text or image is required"), "invalid request not image or text given", utils.Object{"path": "routes/scan.go"})
		c.JSON(http.StatusBadRequest, types.ApiResponse{
			Status:  "error",
			Message: "text or image is required",
		})
		return
	}

	if requestData.AdditionalContext != nil && len(*requestData.AdditionalContext) > 400 {
		utils.LogError(errors.New("additionalContext cannot be too large"), "AdditionalContext > 400", utils.Object{"path": "routes/scan.go", "user": user.ID})
		c.JSON(http.StatusBadRequest, types.ApiResponse{
			Status:  "error",
			Message: "context is too large",
		})
		return
	}

	// We need either text or image
	if requestData.Tone == "" {
		utils.LogError(errors.New("tone is needed it"), "tone is empty", utils.Object{"path": "routes/scan.go", "user": user.ID})
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
		utils.LogError(err, "Error generating responses from OpenAI", utils.Object{"path": "routes/scan.go", "user": user.ID})
		c.JSON(http.StatusInternalServerError, types.ApiResponse{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	if err := retrieveAndIncrementCounts(c, usages); err != nil {
		utils.LogError(err, "Failed to process user counts", utils.Object{"path": "routes/scan.go", "user": user.ID})
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
		return fmt.Errorf("error getting user from context for increment usages: %w", err)
	}

	if err := user.IncrementCountsAndUsages(usages); err != nil {
		return fmt.Errorf("error incrementing user usages count: %w", err)
	}

	return nil
}
