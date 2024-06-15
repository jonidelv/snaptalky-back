package openai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"snaptalky/routes"
)

func GenerateResponses(requestData routes.RequestData) ([]string, error) {
	var promptTemplate string
	switch requestData.Context {
	case "flirting":
		promptTemplate = `Act as the go-to texting assistant for contemporary charm in flirting scenarios, with a repertoire that includes the latest trends, new slang, and a humorous twist. It's programmed to adapt to the user's conversational style, mimicking their mannerisms and preferences to keep the interaction authentic. Be concise and act like an 18-year-old big brother. When responding, assess the user's CONTEXT and TONE provided at the end of this message and incorporate similar language and humor, ensuring a seamless exchange that feels natural. In ambiguous situations, ask for clarifications rather than making assumptions. Your objective is to provide personalized, witty advice for a user trying to interact with a potential dating partner via text, hinge, or tinder. Use this CONTEXT %s and this TONE %s. Respond using this prompt to %s: %s. `
	case "professional":
		promptTemplate = `Act as the go-to texting assistant for contemporary charm in professional interactions, with a repertoire that includes the latest trends, new slang, and a professional twist. It's programmed to adapt to the user's conversational style, mimicking their mannerisms and preferences to keep the interaction authentic. Be concise and act like an 18-year-old big brother. When responding, assess the user's CONTEXT and TONE provided at the end of this message and incorporate similar language and professionalism, ensuring a seamless exchange that feels natural. In ambiguous situations, ask for clarifications rather than making assumptions. Your objective is to provide personalized, professional advice for a user interacting with a boss or colleague via text or apps like LinkedIn or Slack. Use this CONTEXT %s and this TONE %s. Respond using this prompt to %s: %s. `
	default:
		// Case for friendly type as well
		promptTemplate = `Act as the go-to texting assistant for contemporary charm in friendly interactions, with a repertoire that includes the latest trends, new slang, and a humorous twist. It's programmed to adapt to the user's conversational style, mimicking their mannerisms and preferences to keep the interaction authentic. Be concise and act like an 18-year-old big brother. When responding, assess the user's CONTEXT and TONE provided at the end of this message and incorporate similar language and humor, ensuring a seamless exchange that feels natural. In ambiguous situations, ask for clarifications rather than making assumptions. Your objective is to provide personalized, witty advice for a user interacting with a friend via text or apps like Instagram or Snapchat. Use this CONTEXT %s and this TONE %s. Respond using this prompt to %s: %s. `
	}

	comesWithImage := requestData.Image != ""
	inputType, inputContent := "this text message", requestData.Text

	if comesWithImage {
		inputType, inputContent = "the image I will provide below", requestData.Image
	}

	var prompt = fmt.Sprintf(promptTemplate, requestData.Context, requestData.Tone, inputType, inputContent)

	give5AnswersShort := `This is a response to a chat. Make the answers short. Give 5 possible answers/responses in the format specified below. `
	extractLanguage := `Extract the language from the image or text I am passing that needs an answer, and put it in the language key in the response format. `
	translateTones := `I'll specify different tones, you need to translate those tones to the language of the image or the text you are answering to. Here are the different tones that you need to translate and then add to the response: flirtingTones = ["💬 Answer short & crisp","🗣️ Answer long & detailed","❓ Ask a question","❤️ Ask for date","😏 Be cocky & funny","😇 Tease playfully","🤔 Ask intriguing questions","💞 Show affection","🌸 Compliment their style","🌹 Be romantic","🎭 Use role-playing","🕹️ Suggest a fun game","😁 Make them laugh","👻 Be a bit mysterious","🕵️‍♀️ Imagining future together","🍫 Sensual Descriptions","💌 Send a love quote","🎶 Dedicate a song","🌇 Ask about favorite activities","🌠 Ask about their wishes","💫 Make a flirtatious comment","🔍 Dive deeper into a topic","🌙 Wish sweet dreams"]; professionalTones = ["💬 Answer short & crisp","🗣️ Answer long & detailed","❓ Ask a question","👍 Agree with their point","🚫 Disagree respectfully","❓ Ask for clarification","🔄 Change the topic","🗣️ Express your opinion","💬 Paraphrase their point","🎈 Lighten the mood","💼 Stay professional","💬 Start a debate"]; friendlyTones = ["💬 Answer short & crisp","🗣️ Answer long & detailed","❓ Ask a question","😏 Be cocky & funny","😇 Tease playfully","🤔 Ask intriguing questions","👍 Agree with their point","🚫 Disagree respectfully","❓ Ask for clarification","🔄 Change the topic","🗣️ Express your opinion","💬 Paraphrase their point","💞 Show affection","🌸 Compliment their style","😁 Make them laugh","👻 Be a bit mysterious","🎶 Dedicate a song","🌇 Ask about favorite activities","🌠 Ask about their wishes","💫 Make a flirtatious comment","🔍 Dive deeper into a topic","🎈 Lighten the mood"]. `
	respondFormat := `Respond in the following format only (so I can transform this string response into JSON with JSON.parse): {"respondedOk":true,"language":"language","answers":["answer 1","answer 2","answer 3","answer 4","answer 5"],"contexts":{},"tones":{"flirtingTones":["💬 Answer short & crisp","... translated flirtingTones"],"professionalTones":["💬 Answer short & crisp","... translated professionalTones"],"friendlyTones":["💬 Answer short & crisp","... translated friendlyTones"]}} `
	responseNotSuccessful := `If the context of the image or the message I'm passing to respond to is not a message from a chat or not something this AI can respond to as a text or in the context of a chat, respond in the following format only (so I can transform this string response into JSON with JSON.parse): {"respondedOk":true,"language":"language"}`

	prompt = give5AnswersShort + extractLanguage + translateTones + respondFormat + responseNotSuccessful

	content := []map[string]interface{}{
		{
			"type": "text",
			"text": prompt,
		},
	}

	if comesWithImage {
		// Define the image content map
		imageContent := map[string]interface{}{
			"type": "image_url",
			"image_url": map[string]string{
				"url": requestData.Image,
			},
		}
		content = append(content, imageContent)
	}

	// Create the request body map using the content
	reqBody := map[string]interface{}{
		"model":      "gpt-4o-20240513",
		"max_tokens": 100,
		"messages": []map[string]interface{}{
			{
				"role":    "user",
				"content": content,
			},
		},
	}

	reqBodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(reqBodyBytes))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+os.Getenv("OPENAI_API_KEY"))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var respBody map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
		return nil, err
	}

	choices, ok := respBody["choices"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid response format")
	}

	var responses []string
	for _, choice := range choices {
		if msg, ok := choice.(map[string]interface{})["message"].(map[string]interface{})["content"].(string); ok {
			responses = append(responses, msg)
		}
	}

	return responses, nil
}
