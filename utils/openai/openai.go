package openai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"snaptalky/utils"
)

type ApiResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

func CallOpenaiApi(contentPayload []Content) (string, error) {
	reqBody := map[string]interface{}{
		"model":      "gpt-4o-20240513",
		"max_tokens": 200,
		"messages": []map[string]interface{}{
			{
				"role":    "user",
				"content": contentPayload,
			},
		},
	}

	reqBodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		utils.LogError(err, "Error marshaling request body")
		return "", err
	}

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(reqBodyBytes))
	if err != nil {
		utils.LogError(err, "Error creating new request")
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+os.Getenv("OPENAI_API_KEY"))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		utils.LogError(err, "Error making API request")
		return "", err
	}

	defer func() {
		if cErr := resp.Body.Close(); cErr != nil {
			utils.LogError(cErr, "Error closing response body")
		}
	}()

	var apiResponse ApiResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		utils.LogError(err, "Error decoding response")
		return "", err
	}

	if len(apiResponse.Choices) == 0 {
		err := fmt.Errorf("no choices in response")
		utils.LogError(err, "API response validation error")
		return "", err
	}

	return apiResponse.Choices[0].Message.Content, nil
}
