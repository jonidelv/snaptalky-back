package types

import (
	"github.com/google/uuid"
	"snaptalky/models"
)

// ApiResponse Used in any api response/endpoint
type ApiResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// DataToBuildResponses is the data necessary to build the responses
type DataToBuildResponses struct {
	Text    string `json:"text"`
	Context string `json:"context"`
	Tone    string `json:"tone"`
	Image   string `json:"image"` // base64 encoded image
}

type AppUser struct {
	ID                 uuid.UUID                 `json:"id"`
	DeviceID           string                    `json:"device_id"`
	Age                int                       `json:"age,omitempty"`
	Gender             models.Gender             `json:"gender,omitempty"`
	Bio                string                    `json:"bio,omitempty"`
	PublicID           string                    `json:"public_id"`
	IsPremium          bool                      `json:"is_premium"`
	CommunicationStyle models.CommunicationStyle `json:"communication_style"`
	Tone               models.Tone               `json:"tone"`
}
