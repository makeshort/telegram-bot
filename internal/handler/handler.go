package handler

import tgbotapi "github.com/Syfaro/telegram-bot-api"

type Handler struct {
	bot *tgbotapi.BotAPI
}

func New(bot *tgbotapi.BotAPI) *Handler {
	return &Handler{bot: bot}
}

func (h *Handler) HandleStart(u tgbotapi.Update) {

}
