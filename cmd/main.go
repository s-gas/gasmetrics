package main

import (
	"context"
	"log"

	"github.com/s-gas/hauslab/services/gasmetrics/internal/postgres"
	"github.com/s-gas/hauslab/services/gasmetrics/internal/telegram"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	bot, err := telegram.CreateBot()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Authorized on account", bot.Self.UserName)

	conn, err := postgres.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(context.Background())

	updates := telegram.GetUpdates(bot)

	for update := range updates {
		if update.Message != nil {
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, telegram.GenerateResponse(update.Message.Text))
			msg.ReplyToMessageID = update.Message.MessageID
			bot.Send(msg)
		}
	}
}
