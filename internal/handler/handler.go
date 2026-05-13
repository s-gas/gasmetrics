package handler

import (
	"context"
	"errors"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jackc/pgx/v5"
	"github.com/s-gas/hauslab/services/gasmetrics/internal/postgres"
	"github.com/s-gas/hauslab/services/gasmetrics/internal/telegram"
)

func Update(ctx context.Context, conn *pgx.Conn, bot *tgbotapi.BotAPI, update tgbotapi.Update) error {
	if update.Message == nil {
		return errors.New("Update: message is empty")
	}
	lastQuery, err := postgres.GetLastQuery(ctx, conn)
	if err != nil {
		return fmt.Errorf("Update: GetLastQuery: %w", err)
	}
	fmt.Println(lastQuery)
	response := telegram.GetResponse(update.Message.Text)
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, response)
	msg.ReplyToMessageID = update.Message.MessageID
	bot.Send(msg)
	return nil
}
