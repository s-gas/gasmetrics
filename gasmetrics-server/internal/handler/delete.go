package handler

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/s-gas/hauslab/services/gasmetrics/gasmetrics-server/internal/postgres"
)

func Delete(ctx context.Context, conn *pgx.Conn, reading postgres.Reading) error {
	return postgres.DeleteEntry(ctx, conn, reading)
}
