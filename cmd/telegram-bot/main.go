package main

import (
	"log"
	"telegram-bot/internal/bot"
	"telegram-bot/internal/config"
)

func main() {
	cfg := config.MustLoad()

	b, err := bot.New(cfg.Telegram.Token)
	if err != nil {
		log.Printf("can't start bot")
	}

	b.Run()
}
