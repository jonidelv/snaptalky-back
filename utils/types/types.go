package types

import (
	"github.com/google/uuid"
)

// ApiResponse Used in any api response/endpoint
type ApiResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// DataToBuildResponses is the data necessary to build the responses from openai
type DataToBuildResponses struct {
	Text              *string `json:"text"`
	Image             *string `json:"image"`
	Tone              string  `json:"tone" binding:"oneof=flirting friendly formal"`
	ResponseType      *string `json:"responseType"`
	AdditionalContext *string `json:"additionalContext"`
	Location          *string `json:"location"`
	UserBio           *string `json:"userBio"`
	UserGender        *string `json:"userGender"`
	UserAge           *int    `json:"userAge"`
	PreviousResponses *string `json:"previousResponses"`
}

type AppUser struct {
	ID                 uuid.UUID `json:"id"`
	DeviceID           string    `json:"deviceID"`
	Age                int       `json:"age,omitempty"`
	Gender             string    `json:"gender,omitempty"`
	Bio                string    `json:"bio,omitempty"`
	PublicID           string    `json:"publicID"`
	IsPremium          bool      `json:"isPremium"`
	CommunicationStyle string    `json:"communicationStyle"`
	ScanCount          int       `json:"scanCount"`
	Lang               string    `json:"lang"`
}
