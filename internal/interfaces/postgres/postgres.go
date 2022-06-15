package postgres

import (
	"embed"
	"errors"
	"github.com/Nkez/date/internal/model"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

//go:embed migrations/*.sql
var MigrationsFS embed.FS

func NewPostgresDB(cfg *model.Config) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", cfg.DBUrl)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func Up(cfg *model.Config) error {
	source, err := iofs.New(MigrationsFS, "migrations")
	if err != nil {
		return err
	}
	instance, err := migrate.NewWithSourceInstance("iofs", source, cfg.DBUrl)
	if err != nil {
		return err
	}
	if err := instance.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			return nil
		}
		return err
	}
	return nil
}

func Down(cfg *model.Config) error {
	source, err := iofs.New(MigrationsFS, "migrations")
	if err != nil {
		return err
	}
	instance, err := migrate.NewWithSourceInstance("iofs", source, cfg.DBUrl)
	if err != nil {
		return err
	}
	if err := instance.Down(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			return nil
		}
		return err
	}
	return nil
	return nil
}
