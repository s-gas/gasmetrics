package handler

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/s-gas/hauslab/services/gasmetrics/gasmetrics-server/internal/postgres"
)

func StoreValue(ctx context.Context, conn *pgx.Conn, value string) error {
	if value.Message == nil {
		return errors.New("Message is empty")
	}
	lastEntry, err := postgres.GetLastEntry(ctx, conn)
	if err != nil {
		return err
	}
	entry, err := validate(value, lastEntry)
	if err != nil {
		return err
	}
	err = postgres.AddEntry(ctx, conn, entry)
	if err != nil {
		return err
	}
	return nil
}
