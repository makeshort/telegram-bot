package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"io"
	"log/slog"
	"net/http"
	"regexp"
	"telegram-bot/internal/lib/logger/sl"
)

type CreatedURL struct {
	URL   string `json:"url"`
	Alias string `json:"alias"`
}

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
	fmt.Print("start")
}

func (h *Handler) HandleMessage(u tgbotapi.Update) {
	var text = u.Message.Text

	var message tgbotapi.MessageConfig
	if isValidURL(text) {
		body := []byte(fmt.Sprintf(`{"url": "%s"}`, text))
		req, err := http.NewRequest("POST", "https://sh.jus1d.ru/api/url", bytes.NewBuffer(body))
		if err != nil {
			h.log.Error("Error sending request", sl.Err(err))
			return
		}

		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}

		response, err := client.Do(req)
		if err != nil {
			h.log.Error("Error sending request", sl.Err(err))
			return
		}
		defer response.Body.Close()

		if response.StatusCode != 201 {
			h.log.Error("Error while creating URL", slog.Int("response_code", response.StatusCode))
			return
		}

		responseBody, err := io.ReadAll(response.Body)
		if err != nil {
			h.log.Error("Error reading response", sl.Err(err))
			return
		}

		var createdUrl CreatedURL
		err = json.Unmarshal(responseBody, &createdUrl)
		if err != nil {
			h.log.Error("Error parsing JSON:", sl.Err(err))
			return
		}

		h.log.Info("URL created", slog.String("url", text), slog.String("shorter", "https://sh.jus1d.ru/s/"+createdUrl.Alias))

		message = tgbotapi.NewMessage(u.Message.Chat.ID, "**Your shorted URL:** https://sh.jus1d.ru/s/"+createdUrl.Alias)
	} else {
		message = tgbotapi.NewMessage(u.Message.Chat.ID, "**Typed URL is invalid. Try anther one**")
	}

	message.ParseMode = tgbotapi.ModeMarkdown

	_, err := h.bot.Send(message)
	if err != nil {
		h.log.Error("failed to send message", sl.Err(err))
	}
}

func isValidURL(url string) bool {
	urlPattern := `^(https?|ftp):\/\/[^\s\/$.?#].[^\s]*$`

	match, _ := regexp.MatchString(urlPattern, url)
	return match
}
