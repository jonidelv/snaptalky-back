package openai

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"log"
	"os"
	"snaptalky/utils"
)

type ApiResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
}

func CallOpenaiApi(contentPayload []Content) (string, error) {
	reqBody := map[string]interface{}{
		"model":      "chatgpt-4o-latest",
		"max_tokens": 1000,
		"messages": []map[string]interface{}{
			{
				"role":    "user",
				"content": contentPayload,
			},
		},
	}

	client := resty.New()
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer "+os.Getenv("OPENAI_API_KEY")).
		SetBody(reqBody).
		SetResult(&ApiResponse{}).
		Post("https://api.openai.com/v1/chat/completions")

	if err != nil {
		utils.LogError(err, "Error making API request")
		return "", err
	}

	apiResponse := resp.Result().(*ApiResponse)

	if len(apiResponse.Choices) == 0 {
		err := fmt.Errorf("no choices in response")
		utils.LogError(err, "API response validation error")
		return "", err
	}

	log.Printf("Prompt Tokens: %d, Completion Tokens: %d, Total Tokens: %d",
		apiResponse.Usage.PromptTokens, apiResponse.Usage.CompletionTokens, apiResponse.Usage.TotalTokens)

	fmt.Println(apiResponse.Choices[0].Message.Content)

	return apiResponse.Choices[0].Message.Content, nil
}
