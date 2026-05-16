package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func GetLastEntry(ctx context.Context, conn *pgx.Conn) (int, error) {
	var last int
	err := conn.QueryRow(ctx, `
		SELECT value FROM gasmetrics ORDER BY recorded_at DESC LIMIT 1
	`).Scan(&last)
	if err == pgx.ErrNoRows {
		last = 0
	} else if err != nil {
		return 0, fmt.Errorf("GetLastQuery: %w", err)
	}
	return last, nil
}
