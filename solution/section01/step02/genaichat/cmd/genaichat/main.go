package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	"github.com/gohandson/genai-ja/solution/section01/step02/genaichat"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	if len(os.Args) < 1 {
		return errors.New("名前を指定してください。")
	}
	name := os.Args[1]
	bot := genaichat.NewBot(name)

	fmt.Printf("%s: %s\n", bot.Name, bot.FirstMessage)
	fmt.Print("> ")

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		resp := bot.Send(s.Text())

		fmt.Printf("%s: %s\n", bot.Name, resp)
		fmt.Print("> ")
	}

	return nil
}
