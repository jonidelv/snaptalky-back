package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"snaptalky/utils/openai"
)

type RequestData struct {
	Text    string `json:"text"`
	Context string `json:"context"`
	Tone    string `json:"tone"`
	Image   string `json:"image"` // base64 encoded image
}

func ProcessResponse(c *gin.Context) {
	var requestData RequestData

	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Generate responses using the OpenAI API
	responses, err := openai.GenerateResponses(requestData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"responses": responses})
}
