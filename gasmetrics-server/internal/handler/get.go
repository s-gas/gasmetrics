package handler

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/s-gas/hauslab/services/gasmetrics/gasmetrics-server/internal/postgres"
)

func Get(ctx context.Context, conn *pgx.Conn) (int, error) {
	if lastEntry, err := postgres.GetLastEntry(ctx, conn); err != nil {
		return 0, err
	} else {
		return lastEntry, nil
	}
}
