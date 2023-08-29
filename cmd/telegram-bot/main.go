package main

import (
	"telegram-bot/internal/bot"
	"telegram-bot/internal/config"
)

func main() {
	cfg := config.MustLoad()

	b, err := bot.New(cfg)
	if err != nil {
		panic("bot: failed to start bot")
	}

	b.Run()
}
