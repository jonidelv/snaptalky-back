package openai

import (
	"encoding/json"
	"snaptalky/utils"
	"snaptalky/utils/types"
)

type Response struct {
	RespondedOk bool     `json:"respondedOk"`
	Responses   []string `json:"responses"`
	Suggestions []string `json:"suggestions"`
}

func GenerateResponses(dataToBuildResponse *types.DataToBuildResponses) (Response, error) {
	contentPayload := MakeOpenaiContentPayload(dataToBuildResponse)
	openaiResponse, err := CallOpenaiApi(contentPayload)
	if err != nil {
		utils.LogError(err, "Error calling OpenAI API")
		return Response{}, err
	}

	var apiResponse Response
	err = json.Unmarshal([]byte(openaiResponse), &apiResponse)
	if err != nil {
		utils.LogError(err, "Error parsing JSON response")
		return Response{}, err
	}

	return apiResponse, nil
}
