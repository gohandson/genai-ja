package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"os"

	"github.com/gohandson/genai-ja/solution/section02/step04/genaichat"
	"github.com/gohandson/genai-ja/solution/section02/step04/genaichat/server"
)

var (
	flagName string
	flagPort string
)

func init() {
	flag.StringVar(&flagName, "name", "gopher", "name of chat bot")
	flag.StringVar(&flagPort, "port", "8080", "port")
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

	addr := net.JoinHostPort("", flagPort)
	s := server.New(addr, bot)

	return s.ListenAndServe()
}
