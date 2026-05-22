package postgres

import (
	"context"
	"fmt"
	"time"
	"errors"
	"github.com/jackc/pgx/v5"
)

type Reading struct {
	Id		int				`json:"id"`
	Value int 			`json:"value"`
	Date 	time.Time `json:"recorded_at"`
}

func GetPreviousEntry(ctx context.Context, conn *pgx.Conn, reading Reading) (int, error) {
	var last int
	err := conn.QueryRow(ctx, `
		SELECT value FROM gasmetrics 
		WHERE recorded_at < $1 ORDER BY recorded_at DESC LIMIT 1
	`, reading.Date).Scan(&last)
	if err == pgx.ErrNoRows {
		last = 0
	} else if err != nil {
		return 0, fmt.Errorf("GetLastQuery: %w", err)
	}
	return last, nil
}

func GetNextEntry(ctx context.Context, conn *pgx.Conn, reading Reading) (int, error) {
	var next int
	err := conn.QueryRow(ctx, `
		SELECT value FROM gasmetrics
		WHERE recorded_at > $1 ORDER BY recorded_at ASC LIMIT 1
	`, reading.Date).Scan(&next)
	if err == pgx.ErrNoRows {
		next = 0
	} else if err != nil {
		return 0, fmt.Errorf("GetLastQuery: %w", err)
	}
	return next, nil
}

func GetEntries(ctx context.Context, conn *pgx.Conn, limit int) ([]Reading, error) {
	if limit <= 0 {
		return nil, errors.New("GetEntries: limit must be positive")
	}
	query := fmt.Sprintf("SELECT * FROM gasmetrics ORDER BY recorded_at DESC LIMIT %v", limit)
	rows, err := conn.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("GetEntries: %w", err)
	}
	defer rows.Close()

	var entries []Reading
	for rows.Next() {
		var r Reading
		if err := rows.Scan(&r.Id, &r.Value, &r.Date); err != nil {
        return nil, fmt.Errorf("GetEntries scan: %w", err)
    }
		entries = append(entries, r)
	}
	return entries, nil
}
