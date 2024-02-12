package database

import (
	"database/sql"

	_ "embed"
)

var (
	//go:embed query/migration/up.sql
	migrationUpQuery string

	//go:embed query/migration/down.sql
	migrationDownQuery string
)

type Migration struct {
	db *sql.DB
}

func NewMigration(db *sql.DB) Migration {
	return Migration{
		db: db,
	}
}

func (m Migration) Up() error {
	_, err := m.db.Exec(migrationUpQuery)
	return err
}

func (m Migration) Down() error {
	_, err := m.db.Exec(migrationDownQuery)
	return err
}
