package main

import (
	"fmt"
	"os"
)

func generateUrl() (string, error) {
	user := os.Getenv("POSTGRES_USER")
	password, err := os.ReadFile("/run/secrets/postgres_password")
	if err != nil {
		return "", fmt.Errorf("generateUrl: %w\n", err)
	}
	db := os.Getenv("POSTGRES_DB")
	url := fmt.Sprintf("postgres://%s:%s@postgres:5432/%s", user, password, db)
	return url, nil
}
