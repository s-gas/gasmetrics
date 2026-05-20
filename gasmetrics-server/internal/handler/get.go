package handler

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/s-gas/hauslab/services/gasmetrics/gasmetrics-server/internal/postgres"
)


func Get(ctx context.Context, conn *pgx.Conn, limit int) ([]postgres.Reading, error) {
	if entries, err := postgres.GetEntries(ctx, conn, limit); err != nil {
		return nil, err
	} else {
		return entries, nil
	}
}
