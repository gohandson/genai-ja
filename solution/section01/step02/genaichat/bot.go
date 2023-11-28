package genaichat

import (
	"fmt"
)

type Bot struct {
	FirstMessage string
	Name         string
}

func NewBot(name string) *Bot {
	msg := fmt.Sprintf("こんにちは。%sです。オウム返しします。", name)

	return &Bot{
		FirstMessage: msg,
		Name:         name,
	}
}

func (b *Bot) Send(msg string) string {
	return msg
}
