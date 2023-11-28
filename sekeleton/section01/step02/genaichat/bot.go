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

// TODO: 引数でもらった文字列をそのまま返すSendメソッドを作成する
