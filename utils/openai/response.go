package openai

import (
	"encoding/json"
	"github.com/jonidelv/snaptalky-back/utils"
	"github.com/jonidelv/snaptalky-back/utils/types"
)

type Response struct {
	RespondedOk bool     `json:"respondedOk"`
	Responses   []string `json:"responses"`
}

func GenerateResponses(dataToBuildResponse *types.DataToBuildResponses) (Response, int, error) {
	contentPayload := MakeOpenaiContentPayload(dataToBuildResponse)
	var usages int
	openaiResponse, usages, err := CallOpenaiApi(contentPayload)
	if err != nil {
		utils.LogError(err, "Error calling OpenAI API")
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
