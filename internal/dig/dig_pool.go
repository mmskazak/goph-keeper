package dig

import (
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)

func RegisterPgxPool(pool *pgxpool.Pool) error {
	if err := container.Provide(pool); err != nil {
		return fmt.Errorf("error creating dependency graph: %w", err)
	}
	return nil
}

func GetPgxPool() (*pgxpool.Config, error) {
	var pool *pgxpool.Config
	if err := container.Invoke(&pool); err != nil {
		return nil, fmt.Errorf("error invoking dependencies: %w", err)
	}
	return pool, nil
}
