package handler

import (
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"log/slog"
)

type Handler struct {
	log *slog.Logger
	bot *tgbotapi.BotAPI
}

func New(log *slog.Logger, bot *tgbotapi.BotAPI) *Handler {
	return &Handler{
		log: log,
		bot: bot,
	}
}

func (h *Handler) HandleStart(u tgbotapi.Update) {

}
