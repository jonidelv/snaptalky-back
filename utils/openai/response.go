package openai

import (
	"encoding/json"
	"github.com/jonidelv/snaptalky-back/utils"
	"github.com/jonidelv/snaptalky-back/utils/types"
	"time"
)

type Response struct {
	RespondedOk bool     `json:"respondedOk"`
	Responses   []string `json:"responses"`
}

func GenerateResponses(dataToBuildResponse *types.DataToBuildResponses) (Response, int, error) {
	contentPayload := MakeOpenaiContentPayload(dataToBuildResponse)

	var openaiResponse string
	var usages int
	var err error

	// Retry logic for CallOpenaiApi
	for i := 0; i < 2; i++ {
		openaiResponse, usages, err = CallOpenaiApi(contentPayload)
		if err == nil {
			break
		}

		time.Sleep(400 * time.Millisecond)
	}

	// If still failing after retries, return the error
	if err != nil {
		utils.LogError(err, "Error calling OpenAI API after retries")
		return Response{}, 0, err
	}

	var apiResponse Response
	err = json.Unmarshal([]byte(openaiResponse), &apiResponse)
	if err != nil {
		utils.LogError(err, "Error parsing JSON response")
		return Response{}, 0, err
	}

	return apiResponse, usages, nil
}
