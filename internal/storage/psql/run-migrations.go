package psql

import (
	"embed"
	"errors"
	"fmt"
	"gophKeeper/internal/config"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

//go:embed migrations/*
var embedMigrations embed.FS

const migrationsDir = "migrations"

// RunMigrations запускает миграции базы данных из встраиваемой файловой системы.
func RunMigrations(cfg *config.Config) error {
	dir, err := iofs.New(embedMigrations, migrationsDir)
	if err != nil {
		return fmt.Errorf("error opening migrations directory: %w", err)
	}

	m, err := migrate.NewWithSourceInstance("iofs", dir, cfg.DataBaseDSN)
	if err != nil {
		return fmt.Errorf("error opening migrations directory: %w", err)
	}

	if err := m.Up(); err != nil {
		if !errors.Is(err, migrate.ErrNoChange) {
			return fmt.Errorf("error running migrations: %w", err)
		}
	}
	return nil
}
