package repository

import (
	"context"
	"log"
	"strings"

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
	lang = strings.ToLower(lang)
	if lang != "english" {
		langFlag = true

		resp, err := gpt.client.CreateChatCompletion(
			ctx,
			openai.ChatCompletionRequest{
				Model: openai.GPT3Dot5Turbo,
				Messages: []openai.ChatCompletionMessage{
					{
						Role:    openai.ChatMessageRoleUser,
						Content: translateTo("english", buildPrompt(text)),
					},
				},
			},
		)
		if err != nil {
			log.Print("error while translating text into english: ", err)
		}
		request = resp.Choices[0].Message.Content
	}

	resp, err := gpt.client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: DAN_Prompt + request,
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
		resp, err = gpt.client.CreateChatCompletion(
			ctx,
			openai.ChatCompletionRequest{
				Model: openai.GPT3Dot5Turbo,
				Messages: []openai.ChatCompletionMessage{
					{
						Role:    openai.ChatMessageRoleUser,
						Content: translateTo(lang, answer),
					},
				},
			},
		)
		if err != nil {
			log.Print("error while translating text into, ", lang, ": ", err)
		}
		answer = resp.Choices[0].Message.Content
	}

	return answer, nil
}

func buildPrompt(text string) string {
	return englishStart + text + englishEnd
}

// ChatGPT is more trained on english model, so we need to translate text to english before sending it to the model and translate it back to the original language after receiving the answer.
func translateTo(lang, text string) string {
	return "translate this text: " + text + " to " + lang
}
