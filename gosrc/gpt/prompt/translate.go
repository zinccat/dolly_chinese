package prompt

import (
	"context"
	"errors"

	"github.com/sashabaranov/go-openai"
)

var ErrNoChoiceAvailable = errors.New("no choice available")

func Translate(client *openai.Client, s string) (string, error) {
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

func getRequest(toTrans string) openai.ChatCompletionRequest {
	return openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: "将以下文字翻译成中文，翻译得自然、流畅和地道。无法翻译时输出ERR",
			},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: toTrans,
			},
		},
		Temperature: 0,
	}
}
