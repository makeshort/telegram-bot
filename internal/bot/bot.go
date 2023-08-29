package bot

import (
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"log/slog"
	"os"
	"os/signal"
	"reflect"
	"syscall"
	"telegram-bot/internal/config"
	"telegram-bot/internal/handler"
	"telegram-bot/internal/lib/logger/prettyslog"
)

type Bot struct {
	log     *slog.Logger
	client  *tgbotapi.BotAPI
	handler *handler.Handler
}

func New(cfg *config.Config) (*Bot, error) {
	log := initLogger(cfg.Env)
	bot, err := tgbotapi.NewBotAPI(cfg.Telegram.Token)
	if err != nil {
		return nil, err
	}
	h := handler.New(log, bot)
	return &Bot{
		log:     log,
		client:  bot,
		handler: h,
	}, nil
}

func (b *Bot) Run() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	go b.handleUpdates(u)

	b.log.Info("bot successfully started")

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

func initLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case config.EnvLocal:
		log = initPrettyLogger()
	case config.EnvDevelopment:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case config.EnvProduction:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return log
}

func initPrettyLogger() *slog.Logger {
	opts := prettyslog.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	prettyHandler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(prettyHandler)
}
