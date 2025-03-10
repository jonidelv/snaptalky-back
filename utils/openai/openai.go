package openai

import (
	"errors"
	"os"

	"github.com/go-resty/resty/v2"
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

func CallOpenaiApi(contentPayload []Content) (string, int, error) {
	reqBody := map[string]interface{}{
		"model":      "chatgpt-4o-latest",
		"max_tokens": 5000,
		"messages": []map[string]interface{}{
			{
				"role":    "user",
				"content": contentPayload,
			},
		},
		"response_format": map[string]interface{}{
			"type": "json_object",
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
		return "", 0, err
	}

	apiResponse := resp.Result().(*ApiResponse)

	if len(apiResponse.Choices) == 0 {
		err := errors.New("no choices in response")
		return "", 0, err
	}

	//log.Printf("Prompt Tokens: %d, Completion Tokens: %d, Total Tokens: %d",
	//	apiResponse.Usage.PromptTokens, apiResponse.Usage.CompletionTokens, apiResponse.Usage.TotalTokens)

	totalTokens := apiResponse.Usage.TotalTokens

	return apiResponse.Choices[0].Message.Content, totalTokens, nil
}
