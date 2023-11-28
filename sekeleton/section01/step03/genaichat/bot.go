package genaichat

import (
	"context"
	"fmt"

	"github.com/sashabaranov/go-openai"
)

type Bot struct {
	FirstMessage string
	Name         string
	openai       *openai.Client
	req          openai.ChatCompletionRequest
}

func NewBot(ctx context.Context, name, openaiAPIKey string) (*Bot, error) {

	content := fmt.Sprintf("%sという名前のチャットボットとして振る舞ってください。", name)

	req := openai.ChatCompletionRequest{
		Model: openai.GPT4TurboPreview,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: content,
			},
		},
	}

	client := openai.NewClient(openaiAPIKey)

	resp, err := client.CreateChatCompletion(ctx, req)
	if err != nil {
		return nil, err
	}

	return &Bot{
		FirstMessage: resp.Choices[0].Message.Content,
		Name:         name,
		openai:       client,
		req:          req,
	}, nil
}

func (b *Bot) Send(ctx context.Context, msg string) (string, error) {
	b.req.Messages = append(b.req.Messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: msg,
	})

	// TODO: OpenAIのAPI CreateChatCompletionを使ってGPT4にRequestを送ってください
	// ヒント: NewBotで同じことをやっています。エラー処理に注意。

	return resp.Choices[0].Message.Content, nil
}
