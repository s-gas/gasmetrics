package telegram

import (
	"fmt"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func CreateBot() (*tgbotapi.BotAPI, error) {
	token, err := os.ReadFile("/run/secrets/telegram_token")
	if err != nil {
		return nil, fmt.Errorf("CreateBot: %w", err)
	}
	bot, err := tgbotapi.NewBotAPI(string(token))
	if err != nil {
		return nil, fmt.Errorf("CreateBot: %w", err)
	}
	return bot, nil
}
