package database

import (
	"database/sql"

	"github.com/gustapinto/go_hex/internal/core/entity"
	"github.com/gustapinto/go_hex/internal/core/repository"

	_ "embed"
)

var (
	// Validate interface compliance on compile time
	_ repository.Account = (*Account)(nil)

	//go:embed query/account/get_by_id.sql
	accountGetByIDQuery string

	//go:embed query/account/get_all.sql
	accountGetAllQuery string

	//go:embed query/account/create.sql
	accountCreateQuery string

	//go:embed query/account/update_by_id.sql
	accountUpdateByIDQuery string

	//go:embed query/account/delete_by_id.sql
	accountDeleteByIDQuery string
)

type Account struct {
	db *sql.DB
}

func NewAccount(db *sql.DB) Account {
	return Account{
		db: db,
	}
}

func (ac Account) GetByID(id int64) (account entity.Account, err error) {
	row := ac.db.QueryRow(accountGetByIDQuery, id)

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

func (ac Account) GetAll() (accounts []entity.Account, err error) {
	rows, err := ac.db.Query(accountGetAllQuery)
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

func (ac Account) Create(name string, initialValue float64) (id int64, err error) {
	row := ac.db.QueryRow(accountCreateQuery, name, initialValue, initialValue)
	if row.Err() != nil {
		return 0, row.Err()
	}

	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return
}

func (ac Account) UpdateByID(id int64, name string, currentValue float64) error {
	_, err := ac.db.Exec(accountUpdateByIDQuery, name, currentValue, id)
	return err
}

func (ac Account) DeleteByID(id int64) error {
	_, err := ac.db.Exec(accountDeleteByIDQuery, id)
	return err
}
