package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func AddEntry(ctx context.Context, conn *pgx.Conn, reading Reading) error {
	_, err := conn.Exec(ctx, `
		INSERT INTO gasmetrics (value, recorded_at) VALUES ($1, $2)
	`, reading.Value, reading.Date)
	if err != nil {
		return fmt.Errorf("addEntry: %w", err)
	}
	return nil
}
