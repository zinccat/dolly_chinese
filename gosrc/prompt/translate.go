package prompt

import (
	"context"
	"errors"
	"github.com/sashabaranov/go-openai"
)

var ErrNoChoiceAvailable = errors.New("no choice available")

func Translate(s string) (string, error) {
	ctx := context.Background()
	resp, err := client.CreateChatCompletion(ctx, getRequest(s))
	if err != nil {
		return "", err
	}
	if len(resp.Choices) == 0 {
		return "", ErrNoChoiceAvailable
	}
	return resp.Choices[0].Message.Content, nil
}

var client *openai.Client

func getRequest(toTrans string) openai.ChatCompletionRequest {
	return openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: "将以下文字翻译成中文，翻译得自然、流畅和地道",
			},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: toTrans,
			},
		},
		Temperature: 0,
	}
}
