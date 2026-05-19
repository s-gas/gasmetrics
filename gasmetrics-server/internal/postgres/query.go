package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type Reading struct {
	Value int `json:"value"`
}

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

func GetEntries(ctx context.Context, conn *pgx.Conn) ([]Reading, error) {
	rows, err := conn.Query(ctx, `
		SELECT value FROM gasmetrics ORDER BY recorded_at DESC
	`)
	if err != nil {
		return nil, fmt.Errorf("GetEntries: %w", err)
	}
	defer rows.Close()

	var entries []Reading
	for rows.Next() {
		var r Reading
		if err := rows.Scan(&r.Value); err != nil {
        return nil, fmt.Errorf("GetEntries scan: %w", err)
    }
		entries = append(entries, r)
	}
	return entries, nil
}
