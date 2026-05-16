package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func AddEntry(ctx context.Context, conn *pgx.Conn, value int) error {
	_, err := conn.Exec(ctx, `
		INSERT INTO gasmetrics (value) VALUES ($1)
	`, value)
	if err != nil {
		return fmt.Errorf("addEntry: %w", err)
	}
	return nil
}
