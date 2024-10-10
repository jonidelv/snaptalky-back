package openai

import (
	"fmt"
	"snaptalky/utils/types"
)

type TextContent struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

// NewTextContent creates a new TextContent with type set to "text"
func NewTextContent(text string) TextContent {
	return TextContent{
		Type: "text",
		Text: text,
	}
}

func (t TextContent) getOpenaiContentType() string {
	return t.Type
}

type ImageContent struct {
	Type     string            `json:"type"`
	ImageURL map[string]string `json:"image_url"`
}

// NewImageContent creates a new ImageContent with type set to "image_url"
func NewImageContent(imageURL string) ImageContent {
	return ImageContent{
		Type: "image_url",
		ImageURL: map[string]string{
			"url": imageURL,
		},
	}
}

func (i ImageContent) getOpenaiContentType() string {
	return i.Type
}

type Content interface {
	getOpenaiContentType() string
}

func MakeOpenaiContentPayload(data *types.DataToBuildResponses) []Content {
	// TEMPLATES
	// Tone
	var toneTemplate string
	switch data.Tone {
	case "flirting":
		toneTemplate = `Act as the go-to texting assistant for contemporary charm in flirting scenarios, with a 
repertoire that includes the latest trends, new slang, and a humorous twist. Be concise and respond short, 
not acting like an AI but rather as a human responding to a message. Your objective is to provide 
personalized, witty advice for a user trying to interact with a potential dating partner via text, hinge, 
or tinder.`
	case "formal":
		toneTemplate = `Act as the go-to texting assistant for contemporary charm in professional interactions, 
with a repertoire that includes the latest trends, new slang, and a professional twist. Be concise and 
respond short, not acting like an AI but rather as a human responding to a message. Your objective is to 
provide personalized, professional or formal advice for a user interacting with a boss, colleague or friend 
in a formal way via text or apps like LinkedIn or Slack.`
	case "friendly":
		fallthrough
	default:
		toneTemplate = `Act as the go-to texting assistant for contemporary charm in friendly interactions, with 
a repertoire that includes the latest trends, new slang, and a humorous twist. Be concise and respond 
short, not acting like an AI but rather as a human responding to a message. Your objective is to provide 
personalized, witty advice for a user interacting with a friend via text or apps like Instagram or 
Snapchat. Respond in a friendly way.`
	}

	// Additional context
	var additionalContextTemplate string
	hasText := data.Text != nil && *data.Text != ""
	hasContextText := data.ContextText != nil && *data.ContextText != ""
	prefix := `This is additional context passed by the user to use when building the responses to this message, use 
this information only if you think provide additional information that can be useful in order to build the response.`
	switch {
	case hasText && hasContextText:
		additionalContextTemplate = fmt.Sprintf(`%s %s %s.`, prefix, *data.Text, *data.ContextText)
	case hasText:
		additionalContextTemplate = fmt.Sprintf(`%s %s.`, prefix, *data.Text)
	case hasContextText:
		additionalContextTemplate = fmt.Sprintf(`%s %s.`, prefix, *data.ContextText)
	default:
		additionalContextTemplate = ""
	}

	//Location
	var locationTemplate string
	hasLocation := data.Location != nil && *data.Location != ""
	switch {
	case hasLocation:
		locationTemplate = fmt.Sprintf(`Use this location: %s to buil the answers in a way someone from this locaton 
would respond, use the slangs and modisms fromt the location provided.`, *data.Location)
	default:
		locationTemplate = ""
	}

	// Context
	var contextTemplate string
	hasContext := data.Context != nil && *data.Context != ""
	prefix = `THE GENERATED RESPONSES NEED TO BE ONLY IN THIS WAY!: `
	switch {
	case hasContext:
		contextTemplate = fmt.Sprintf(`%s %s.`, prefix, *data.Context)
	default:
		contextTemplate = ""
	}

	// User input
	var userInputTemplate string
	hasImage := data.Image != nil && *data.Image != ""
	switch {
	case hasImage:
		userInputTemplate = "The responses to be generated needs to be based on the image passed as base64"
	default:
		userInputTemplate = fmt.Sprintf("The responses to be generated needs to be a response to this USER INPUT -> %s", *data.Text)
	}

	templates := fmt.Sprintf("%s %s %s %s %s.", toneTemplate, additionalContextTemplate, locationTemplate, contextTemplate, userInputTemplate)

	//FORMATS
	// 8 Responses
	give8AnswersFormat := `This is a response to a chat. Make the answers short. Give 8 possible answers/responses in the 
format specified next: `

	// Format
	respondFormat := `Respond in the following format only (so I can transform this string response into JSON with JSON.parse): 
{"respondedOk":true,"responses":["responses 1","responses 2","responses 3","responses 4","responses 5", "responses 6", "responses 7", "responses 8"]}`

	// Language
	languageFormart := `Extract the language from the image or USER INPUT I am passing that needs an answer to, and use it 
to RESPOND WITH THE RESPONSES!`

	// Input is not a chat
	inputNotChatFormat := `If the image or the USER INPUT I'm passing to respond to is not a message from 
a chat or not something this AI can respond to as a text or in the context of a chat, respond in the following format 
only (so I can transform this string response into JSON with JSON.parse): {"respondedOk":false}`

	formats := fmt.Sprintf("%s %s %s %s", give8AnswersFormat, respondFormat, languageFormart, inputNotChatFormat)

	prompt := templates + formats

	content := []Content{
		NewTextContent(prompt),
	}

	if hasImage {
		content = append(content, NewImageContent(*data.Image))
	}

	return content
}
