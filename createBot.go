package main

import (
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func createBot() (*tgbotapi.BotAPI, error) {
	token, err := os.ReadFile("/run/secrets/telegram_token")
	if err != nil {
		return nil, err
	}
	bot, err := tgbotapi.NewBotAPI(string(token))
	if err != nil {
		return nil, err
	}
	return bot, nil
}
