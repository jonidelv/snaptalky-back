package types

// ApiResponse Used in any api response/endpoint
type ApiResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// is the data necesary to build the responses
type DataToBuildResponses struct {
	Text    string `json:"text"`
	Context string `json:"context"`
	Tone    string `json:"tone"`
	Image   string `json:"image"` // base64 encoded image
}
