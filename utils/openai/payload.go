package openai

import (
	"fmt"
	"github.com/jonidelv/snaptalky-back/utils/types"
	"strings"
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

// Helper function to check if a string pointer has data
func hasString(s *string) bool {
	return s != nil && *s != ""
}

// Helper function to check if an int pointer has data
func hasInt(i *int) bool {
	return i != nil && *i != 0
}

// MakeOpenaiContentPayload builds the prompt and returns it as a string
func MakeOpenaiContentPayload(data *types.DataToBuildResponses) []Content {
	var promptBuilder strings.Builder

	// Handle message input (text or image)
	hasMessageText := hasString(data.Text)
	hasMessageImage := hasString(data.Image)

	promptBuilder.WriteString("Please perform the following steps:\n\n")

	if hasMessageText && !hasMessageImage {
		promptBuilder.WriteString("A) Extract and identify the language of this text: '" + *data.Text + "'.\n\n")
	}

	if hasMessageImage && !hasMessageText {
		promptBuilder.WriteString("A) Extract the text from the image provided in base64 format below and identify its language.\n\n")
	}

	promptBuilder.WriteString("B) Using the identified language, generate 10 possible replies to this message in a chat, considering the following guidelines:\n\n")
	// Define the language source variable
	var languageSource string

	if hasMessageText && !hasMessageImage {
		promptBuilder.WriteString("- The message to reply to is: '" + *data.Text + "'.\n")
		promptBuilder.WriteString("- Make sure the replies are in the same language as the language extracted in point A).\n")
		promptBuilder.WriteString("Use this information only if it makes sense in the context of a message, if not respond with {\"respondedOk\":false}.\n\n")
		languageSource = "message to reply to"
	}

	if hasMessageImage && !hasMessageText {
		promptBuilder.WriteString("- The message to reply to is an image provided in base64 format. If the content of the image cannot be determined, respond with {\"respondedOk\":false}.\n")
		promptBuilder.WriteString("- Make sure the replies are in the same language as the language extracted in point A).\n\n")
		languageSource = "content of the image provided in base64"
	}

	if hasMessageText && hasMessageImage {
		promptBuilder.WriteString("- The message to reply to includes both text and an image. Use both as context, but ensure that the replies are in the same language as the image provided.\n\n")
		languageSource = "content of the image provided in base64"
	}

	promptBuilder.WriteString("- The conversation is in a '" + data.Tone + "' context. Possible contexts are friendly, formal, or flirting. Adjust the tone of the replies accordingly.\n\n")

	if hasString(data.AdditionalContext) {
		promptBuilder.WriteString("- Additional context provided by the user to help build the responses: '" + *data.AdditionalContext + "'.\n")
		promptBuilder.WriteString("Use the additional context only if it makes sense in the context of the conversation or chat.\n\n")
	}

	if hasString(data.Location) {
		promptBuilder.WriteString("- Location: '" + *data.Location + "'. Use this location to adjust the style of the responses, incorporating local slang, but ensure that the responses remain in the same language as the: " + languageSource + ".\n")
		promptBuilder.WriteString("- IMPORTANT: Regardless of the location or any other context, ensure that all responses are in the same language as the " + languageSource + " using the language extracted in point A). Do not translate or change the language of the replies based on the location.\n\n")
	}

	var userInfoParts []string
	if hasString(data.UserBio) {
		userInfoParts = append(userInfoParts, "Bio: "+*data.UserBio)
	}
	if hasInt(data.UserAge) {
		userInfoParts = append(userInfoParts, fmt.Sprintf("Age: %d", *data.UserAge))
	}
	if hasString(data.UserGender) {
		userInfoParts = append(userInfoParts, "Gender: "+*data.UserGender)
	}
	if len(userInfoParts) > 0 {
		promptBuilder.WriteString("- User information: " + strings.Join(userInfoParts, ", ") + ".\n")
		promptBuilder.WriteString("This is the information and bio of the user that is requesting help with the messages. Use this to have additional information when crafting the responses.\n")
		promptBuilder.WriteString("This information DOES NOT INTENT to modify the language of the replies, use this just as context.\n\n")
	}

	if hasString(data.PreviousResponses) && !hasString(data.ResponseType) {
		promptBuilder.WriteString("- Previous responses chosen by the user in this context: " + *data.PreviousResponses + ".\n")
		promptBuilder.WriteString("IMPORTANT: Compare the language of these previous responses with the language of the " + languageSource + " extracted in point A).\n")
		promptBuilder.WriteString("If the languages match, use the previous responses to understand the user's preferred style and format.\n")
		promptBuilder.WriteString("If the languages do not match, ignore the previous responses, as we do not want to change the language or idiom of the replies.\n\n")
	}

	promptBuilder.WriteString("- Your task is to generate 10 possible short responses that match the conversation context and the type of response specified, using the user's preferences if available.\n\n")
	promptBuilder.WriteString("- If the message to respond to (user input: image or text) is not suitable for generating responses (e.g., it's not a message from a chat), or if the content cannot be determined, respond with {\"respondedOk\":false}.\n\n")
	promptBuilder.WriteString("- Respond in the following format only (so I can transform this string response into JSON with JSON.parse): {\"respondedOk\":true,\"responses\":[\"response 1\",\"response 2\",\"response 3\",\"response 4\",\"response 5\",\"response 6\",\"response 7\",\"response 8\", \"response 9\", \"response 10\"]}\n\n")
	promptBuilder.WriteString("- Do not include any other text in your response.\n\n")

	var toneTemplate string
	switch strings.ToLower(data.Tone) {
	case "flirting":
		toneTemplate = `- Act as the go-to texting assistant for contemporary charm in flirting scenarios, with a 
repertoire that includes the latest trends, new slang, and a humorous twist. Be concise and respond short, 
not acting like an AI but rather as a human responding to a message. Your objective is to provide 
personalized, witty advice for a user trying to interact with a potential dating partner via text, Hinge, 
or Tinder.`
	case "formal":
		toneTemplate = `- Act as the go-to texting assistant for contemporary charm in professional interactions, 
with a repertoire that includes the latest trends, new slang, and a professional twist. Be concise and 
respond short, not acting like an AI but rather as a human responding to a message. Your objective is to 
provide personalized, professional or formal advice for a user interacting with a boss, colleague, or friend 
in a formal way via text or apps like LinkedIn or Slack.`
	case "friendly":
		fallthrough
	default:
		toneTemplate = `- Act as the go-to texting assistant for contemporary charm in friendly interactions, with 
a repertoire that includes the latest trends, new slang, and a humorous twist. Be concise and respond 
short, not acting like an AI but rather as a human responding to a message. Your objective is to provide 
personalized, witty advice for a user interacting with a friend via text or apps like Instagram or 
Snapchat. Respond in a friendly way.`
	}

	promptBuilder.WriteString(toneTemplate + "\n\n")

	if hasString(data.ResponseType) {
		promptBuilder.WriteString("- IMPORTANT: The user wants the responses to be in THIS specific style: '" + *data.ResponseType + "'.\n")
		promptBuilder.WriteString("All 10 responses need to be in this style. Respond only in this way.\n")
		promptBuilder.WriteString("The response style is: '" + *data.ResponseType + "'. The 10 responses need to be in this specific way.\n")
	}

	if hasMessageImage {
		promptBuilder.WriteString("- IMPORTANT: Regardless of previous responses, user preferences, or any other context, ensure that all replies are responded in the same language as the image provided below same language extracted in point A).\n\n")
	} else {
		promptBuilder.WriteString("- IMPORTANT: Regardless of previous responses, user preferences, or any other context, ensure that all replies are responded in the same language as this text: '" + *data.Text + " same as the language extracted in point A)'.\n\n")
	}

	prompt := promptBuilder.String()

	content := []Content{
		NewTextContent(prompt),
	}

	if hasMessageImage {
		content = append(content, NewImageContent(*data.Image))
	}

	return content
}
