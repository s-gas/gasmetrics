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
		return fmt.Errorf("AddEntry: %w", err)
	}
	return nil
}

func DeleteEntry(ctx context.Context, conn *pgx.Conn, reading Reading) error {
	tag, err := conn.Exec(ctx, `
		DELETE FROM gasmetrics WHERE id = $1
	`, reading.Id)
	if err != nil {
		return fmt.Errorf("DeleteEntry: %w", err)
	}
	if tag.RowsAffected() == 0 {
		return fmt.Errorf("Delete Entry: %w", ErrNotFound)
	}
	return nil
}
