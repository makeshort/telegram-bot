package bot

import (
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"os"
	"os/signal"
	"reflect"
	"syscall"
	"telegram-bot/internal/handler"
)

type Bot struct {
	client  *tgbotapi.BotAPI
	handler *handler.Handler
}

func New(token string) (*Bot, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}
	h := handler.New(bot)
	return &Bot{
		client:  bot,
		handler: h,
	}, nil
}

func (b *Bot) Run() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	go b.handleUpdates(u)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
}

func (b *Bot) handleUpdates(u tgbotapi.UpdateConfig) {
	updates, _ := b.client.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		if reflect.TypeOf(update.Message.Text).Kind() == reflect.String && update.Message.Text != "" {
			switch update.Message.Text {
			case "/start":
				b.handler.HandleStart(update)
			}
		}
	}
}
