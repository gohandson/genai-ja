package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/gohandson/genai-ja/solution/section01/step03/genaichat"
)

var (
	flagName string
)

func init() {
	// TODO: flagNameに-nameで指定した名前が入るようにしてください。
	// デフォルト値はgopherにしてください
}

func main() {
	ctx := context.Background()
	if err := run(ctx); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	apikey := os.Getenv("OPENAI_API_KEY")
	bot, err := genaichat.NewBot(ctx, flagName, apikey)
	if err != nil {
		return err
	}

	fmt.Printf("%s: %s\n", bot.Name, bot.FirstMessage)
	fmt.Print("> ")

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		resp, err := bot.Send(ctx, s.Text())
		if err != nil {
			return err
		}

		fmt.Printf("%s: %s\n", bot.Name, resp)
		fmt.Print("> ")
	}

	return nil
}
