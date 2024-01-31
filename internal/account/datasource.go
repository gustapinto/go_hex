package account

import (
	"database/sql"
)

type SqlDataSource struct {
	DB *sql.DB
}

var _ Repository = (*SqlDataSource)(nil) // Validate interface compliance on compile time

func (ri SqlDataSource) GetByID(id int64) (account Account, err error) {
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
	row := ri.DB.QueryRow(query, id)

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

func (ri SqlDataSource) GetAll() (accounts []Account, err error) {
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
	rows, err := ri.DB.Query(query)
	if err != nil {
		return
	}

	for rows.Next() {
		var account Account
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

func (ri SqlDataSource) Count() (total int64, err error) {
	// TODO
	return
}

func (ri SqlDataSource) Create(name string, initialValue float64) (id int64, err error) {
	// TODO
	return
}

func (ri SqlDataSource) UpdateByID(id int64, name string, currentValue float64) error {
	// TODO
	return nil
}

func (ri SqlDataSource) SumToCurrentValue(id int64, value float64) error {
	// TODO
	return nil
}

func (ri SqlDataSource) SubtractFromCurrentValue(id int64, value float64) error {
	// TODO
	return nil
}

func (ri SqlDataSource) DeleteByID(id int64) error {
	// TODO
	return nil
}
