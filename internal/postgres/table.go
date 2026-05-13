package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func createTable(conn *pgx.Conn, ctx context.Context) error {
	_, err := conn.Exec(ctx, `
		CREATE TABLE IF NOT EXISTS gasmetrics (
			id 				SERIAL PRIMARY KEY,
			value			NUMERIC		NOT NULL,
			recorded_at		TIMESTAMPTZ	NOT NULL	DEFAULT NOW()
		)
	`)
	if err != nil {
		return fmt.Errorf("createTable: Exec: %w", err)
	}
	return nil
}
