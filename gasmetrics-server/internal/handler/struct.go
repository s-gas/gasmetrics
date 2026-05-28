package handler

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type Server struct {
	Ctx  context.Context
	Conn *pgx.Conn
}
