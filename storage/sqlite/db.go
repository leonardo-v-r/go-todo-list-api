package storage

import (
	"database/sql"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	"github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	*sql.DB
}

func New(dsn string) (*DB, error) {
	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return &DB{db}, nil
}

func (db *DB) Migrate() error {
	instance, err := sqlite3.WithInstance(db.DB, &sqlite3.Config{})
	if err != nil {
		return err
	}

	fs, err := (&file.File{}).Open("./storage/migrations")
	if err != nil {
		return err
	}

	m, err := migrate.NewWithInstance("file", fs, "sqlite3", instance)
	if err != nil {
		return err
	}

	if err = m.Down(); err != nil {
		return err
	}

	if err = m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}
	return nil
}
