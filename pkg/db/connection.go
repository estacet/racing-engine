package db

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func New(ctx context.Context) (*pgx.Conn, error) {
	connStr := "postgres://postgres:pass@localhost:5433/racing_engine"
	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
