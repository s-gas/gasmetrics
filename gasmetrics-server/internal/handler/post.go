package handler

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/s-gas/hauslab/services/gasmetrics/gasmetrics-server/internal/postgres"
)

func Post(ctx context.Context, conn *pgx.Conn, reading postgres.Reading) error {
	lastEntry, err := postgres.GetPreviousEntry(ctx, conn, reading)
	if err != nil {
		return err
	}
	nextEntry, err := postgres.GetNextEntry(ctx, conn, reading)
	if err != nil {
		return err
	}
	err = validate(reading.Value, lastEntry, nextEntry)
	if err != nil {
		return err
	}
	return postgres.AddEntry(ctx, conn, reading)
}
