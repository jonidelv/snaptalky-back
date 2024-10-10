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

// DataToBuildResponses is the data necessary to build the responses
type DataToBuildResponses struct {
	Text        *string `json:"text"`
	Image       *string `json:"image"`
	Tone        string  `json:"tone" binding:"oneof=flirting friendly formal"`
	Context     *string `json:"context"`
	ContextText *string `json:"contextText"`
	Location    *string `json:"location"`
}

type AppUser struct {
	ID                 uuid.UUID `json:"id"`
	DeviceID           string    `json:"deviceID"`
	Age                int       `json:"age,omitempty"`
	Gender             string    `json:"gender,omitempty" binding:"oneof=male female other"`
	Bio                string    `json:"bio,omitempty"`
	PublicID           string    `json:"publicID"`
	IsPremium          bool      `json:"isPremium"`
	CommunicationStyle string    `json:"communicationStyle" binding:"oneof=default direct passive"`
}
