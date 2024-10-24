package psql

import (
	"context"
	"fmt"
	"gophKeeper/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPgxPool(ctx context.Context, cfg *config.Config) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(ctx, cfg.DataBaseDSN)
	if err != nil {
		return nil, fmt.Errorf("failed NewPostgres to connect: %w", err)
	}

	err = pool.Ping(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to ping dbshorturl connection: %w", err)
	}
	return pool, nil
}
