package repository

import (
	"context"
	"log"

	"github.com/sashabaranov/go-openai"
)

type gptClient struct {
	client *openai.Client
}

func NewGptClient(token string) *gptClient {
	client := openai.NewClient(token)
	return &gptClient{client: client}
}

func (gpt *gptClient) NewRequest(ctx context.Context, text string, lang string) (string, error) {
	request := text
	langFlag := false

	lang, err := detectLang(text)
	if err != nil {
		return "", err
	}
	if lang != "en" { //check if the lang is not english, cause model's answer is more deep in english.
		langFlag = true

		translatedText, err := translateTo("en", text)
		if err != nil {
			return "", err
		}
		request = translatedText
	}

	resp, err := gpt.client.CreateChatCompletion( //send request to the model
		ctx,
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: DAN_Prompt + buildPrompt(request),
				},
			},
		},
	)
	//TODO API returns different errors, need to handle them properly
	if err != nil {
		log.Printf("ChatCompletion error: %v\n", err)
		return "", err
	}
	answer := resp.Choices[0].Message.Content

	if langFlag {
		translatedText, err := translateTo(lang, answer)
		if err != nil {
			return "", err
		}
		answer = translatedText
	}

	return answer, nil
}

func buildPrompt(text string) string {
	return englishStart + text + englishEnd
}

// // ChatGPT is more trained on english model, so we need to translate text to english before sending it to the model and translate it back to the original language after receiving the answer.
// func translateTo(lang, text string) string {
// 	return "translate this text: " + text + " to " + lang
// }
