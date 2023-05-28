package repository

import (
	"context"
	"log"

	"github.com/Phaseant/DreamAnalyzer/internal/model"
	"github.com/sashabaranov/go-openai"
)

type gptClient struct {
	client *openai.Client
}

func NewGptClient(token string) *gptClient {
	client := openai.NewClient(token)
	return &gptClient{client: client}
}

func (gpt *gptClient) NewRequest(ctx context.Context, text string, lang int) (string, error) {
	resp, err := gpt.client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: DAN_Prompt,
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: buildPrompt(text, lang),
				},
			},
		},
	)
	//TODO API returns different errors, need to handle them properly
	if err != nil {
		log.Printf("ChatCompletion error: %v\n", err)
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}

func buildPrompt(text string, lang int) string {
	switch lang {
	case model.English:
		return EnglishStart + text + EnglishEnd
	case model.Russian:
		return RussianStart + text + RussianEnd
	default:
		return EnglishStart + text + EnglishEnd
	}
}
