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
		telegram.Reply("Failure", bot, update)
		return errors.New("Update: message is empty")
	}
	lastEntry, err := postgres.GetLastEntry(ctx, conn)
	if err != nil {
		telegram.Reply("Failure", bot, update)
		return fmt.Errorf("Update: %w", err)
	}
	value, err := validate(update.Message.Text, lastEntry)
	if err != nil {
		telegram.Reply("Failure", bot, update)
		return fmt.Errorf("Update: %w", err)
	}
	_ = value
	telegram.Reply("Success", bot, update)
	return nil
}
