package database

import (
	"database/sql"

	"github.com/gustapinto/go_hex/internal/core/entity"
	"github.com/gustapinto/go_hex/internal/core/repository"
)

var _ repository.Account = (*Account)(nil) // Validate interface compliance on compile time

type Account struct {
	db *sql.DB
}

func NewAccount(db *sql.DB) Account {
	return Account{
		db: db,
	}
}

func (ri Account) GetByID(id int64) (account entity.Account, err error) {
	query := `
		SELECT
			a.id,
			a.name,
			a.initial_value,
			a.current_value,
			a.created_at,
			a.updated_at
		FROM
			account a
		WHERE
			a.id = $1
			AND a.deleted_at IS NULL
	`
	row := ri.db.QueryRow(query, id)

	if row.Err() != nil {
		err = row.Err()
		return
	}

	err = row.Scan(
		&account.ID,
		&account.Name,
		&account.InitialValue,
		&account.CurrentValue,
		&account.CreatedAt,
		&account.UpdatedAt)
	return
}

func (ri Account) GetAll() (accounts []entity.Account, err error) {
	query := `
		SELECT
			a.id,
			a.name,
			a.initial_value,
			a.current_value,
			a.created_at,
			a.updated_at
		FROM
			account a
		WHERE
			a.deleted_at IS NULL
	`
	rows, err := ri.db.Query(query)
	if err != nil {
		return
	}

	for rows.Next() {
		var account entity.Account
		err = rows.Scan(
			&account.ID,
			&account.Name,
			&account.InitialValue,
			&account.CurrentValue,
			&account.CreatedAt,
			&account.UpdatedAt)
		if err != nil {
			return
		}

		accounts = append(accounts, account)
	}
	return
}

func (ri Account) Create(name string, initialValue float64) (id int64, err error) {
	// TODO
	return
}

func (ri Account) UpdateByID(id int64, name string, currentValue float64) error {
	// TODO
	return nil
}

func (ri Account) SumToCurrentValueByID(id int64, value float64) error {
	// TODO
	return nil
}

func (ri Account) SubtractFromCurrentValueByID(id int64, value float64) error {
	// TODO
	return nil
}

func (ri Account) DeleteByID(id int64) error {
	// TODO
	return nil
}
