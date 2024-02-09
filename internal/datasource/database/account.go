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
			id,
			name,
			initial_value,
			current_value,
			created_at,
			updated_at
		FROM
			account a
		WHERE
			id = $1
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
			id,
			name,
			initial_value,
			current_value,
			created_at,
			updated_at
		FROM
			account a
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
	query := `
		INSERT INTO account (
			name,
			initial_value,
			current_value,
			created_at,
			updated_at
		)
		VALUES (
			?,
			?,
			?,
			CURRENT_TIMESTAMP,
			CURRENT_TIMESTAMP
		)
		RETURNING id
	`
	row := ri.db.QueryRow(query, name, initialValue, initialValue)
	if row.Err() != nil {
		return 0, row.Err()
	}

	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return
}

func (ri Account) UpdateByID(id int64, name string, currentValue float64) error {
	query := `
		UPDATE
			account
		SET
			name = ?,
			current_value = ?
		WHERE
			id = ?
	`
	_, err := ri.db.Exec(query, name, currentValue, id)
	return err
}

func (ri Account) SumToCurrentValueByID(id int64, value float64) error {
	query := `
		UPDATE
			account
		SET
			current_value = (current_value + ?)
		WHERE
			id = ?
	`
	_, err := ri.db.Exec(query, value, id)
	return err
}

func (ri Account) SubtractFromCurrentValueByID(id int64, value float64) error {
	query := `
		UPDATE
			account
		SET
			current_value = (current_value - ?)
		WHERE
			id = ?
	`
	_, err := ri.db.Exec(query, value, id)
	return err
}

func (ri Account) DeleteByID(id int64) error {
	query := `
		DELETE FROM
			account
		WHERE
			id = ?
	`
	_, err := ri.db.Exec(query, id)
	return err
}
