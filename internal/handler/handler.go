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
		telegram.Reply("Failed to add entry", bot, update)
		return errors.New("Update: message is empty")
	}
	lastEntry, err := postgres.GetLastEntry(ctx, conn)
	if err != nil {
		telegram.Reply("Failed to add entry", bot, update)
		return fmt.Errorf("Update: %w", err)
	}
	value, err := validate(update.Message.Text, lastEntry)
	if err != nil {
		telegram.Reply("Failed to add entry", bot, update)
		return fmt.Errorf("Update: %w", err)
	}
	err = postgres.AddEntry(ctx, conn, value)
	if err != nil {
		telegram.Reply("Failed to add entry", bot, update)
		return fmt.Errorf("Update: %w", err)
	}
	telegram.Reply("Entry added successfully", bot, update)
	return nil
}
