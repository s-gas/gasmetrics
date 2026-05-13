package postgres

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

func Connect(ctx context.Context) (*pgx.Conn, error) {
	url, err := generateUrl()
	if err != nil {
		return nil, fmt.Errorf("Connect: generateUrl: %w", err)
	}
	conn, err := pgx.Connect(ctx, url)
	if err != nil {
		return nil, fmt.Errorf("Connect: Connect: %w", err)
	}
	err = createTable(ctx, conn)
	if err != nil {
		return nil, fmt.Errorf("Connect: createTable: %w", err)
	}
	return conn, nil
}

func generateUrl() (string, error) {
	user := os.Getenv("POSTGRES_USER")
	password, err := os.ReadFile("/run/secrets/postgres_password")
	if err != nil {
		return "", fmt.Errorf("generateUrl: ReadFile: %w\n", err)
	}
	db := os.Getenv("POSTGRES_DB")
	url := fmt.Sprintf("postgres://%s:%s@postgres:5432/%s", user, password, db)
	return url, nil
}
