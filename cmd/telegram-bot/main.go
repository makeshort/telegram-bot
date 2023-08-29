package main

import (
	"github.com/makeshort/telegram-bot/internal/bot"
	"github.com/makeshort/telegram-bot/internal/config"
	"log"
)

func main() {
	cfg := config.MustLoad()

	b, err := bot.New(cfg.Telegram.Token)
	if err != nil {
		log.Printf("can't start bot")
	}

	b.Run()
}
