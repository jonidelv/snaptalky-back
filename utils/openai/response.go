package openai

import (
	"encoding/json"
	"time"

	"github.com/jonidelv/snaptalky-back/utils/types"
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
		return Response{}, 0, err
	}

	var apiResponse Response
	err = json.Unmarshal([]byte(openaiResponse), &apiResponse)
	if err != nil {
		return Response{}, 0, err
	}

	return apiResponse, usages, nil
}
