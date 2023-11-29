package genaichat

import (
	"bytes"
	"context"

	"github.com/sashabaranov/go-openai"
)

type Bot struct {
	FirstMessage string
	Name         string
	openai       *openai.Client
	req          openai.ChatCompletionRequest
}

func NewBot(ctx context.Context, name, openaiAPIKey string) (*Bot, error) {

	var content bytes.Buffer
	prompt := NewPrompt(name)
	if err := prompt.Write(&content); err != nil {
		return nil, err
	}

	req := openai.ChatCompletionRequest{
		Model: openai.GPT4TurboPreview,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: content.String(),
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

	resp, err := b.openai.CreateChatCompletion(ctx, b.req)
	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}
