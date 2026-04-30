package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	token, err := os.ReadFile("/run/secrets/telegram_token")
	if err != nil {
		log.Fatal(err)
	}
	bot, err := tgbotapi.NewBotAPI(string(token))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Authorized on account %s", bot.Self.UserName)
}
