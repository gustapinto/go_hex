package database

import (
	"database/sql"

	"github.com/gustapinto/go_hex/internal/core/entity"
	"github.com/gustapinto/go_hex/internal/core/repository"
)

// Validate interface compliance on compile time
var _ repository.Transaction = (*Transaction)(nil)

type Transaction struct {
	db *sql.DB
}

func NewTransaction(db *sql.DB) Transaction {
	return Transaction{
		db: db,
	}
}

func (tr Transaction) GetByIDAndAccountID(id, accountID int64) (transaction entity.Transaction, err error) {
	query := `
		SELECT
			id,
			name,
			account_id,
			value,
			created_at
		FROM
			transaction
		WHERE
			id = $1
			AND account_id = $2
	`
	row := tr.db.QueryRow(query, id, accountID)
	if row.Err() != nil {
		return transaction, row.Err()
	}

	err = row.Scan(
		&transaction.ID,
		&transaction.Name,
		&transaction.AccountID,
		&transaction.Value,
		&transaction.CreatedAt)
	if err != nil {
		return
	}
	return
}

func (tr Transaction) GetAllByAccountID(accountID int64) (transactions []entity.Transaction, err error) {
	query := `
		SELECT
			id,
			name,
			account_id,
			value,
			created_at
		FROM
			transaction
		WHERE
			account_id = $1
	`
	rows, err := tr.db.Query(query, accountID)
	if err != nil {
		return
	}

	for rows.Next() {
		var transaction entity.Transaction

		err = rows.Scan(
			&transaction.ID,
			&transaction.Name,
			&transaction.AccountID,
			&transaction.Value,
			&transaction.CreatedAt)
		if err != nil {
			return
		}

		transactions = append(transactions, transaction)
	}
	return
}

func (tr Transaction) CreateByAccountID(accountID int64, value float64, name string) (id int64, err error) {
	query := `
		INSERT INTO transaction (
			name,
			account_id,
			value,
			created_at
		)
		VALUES (
			$1,
			$2,
			$3,
			CURRENT_TIMESTAMP
		)
		RETURNING id
	`
	row := tr.db.QueryRow(query, name, accountID, value)
	if row.Err() != nil {
		return 0, row.Err()
	}

	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return
}

func (tr Transaction) DeleteByIDAndAccountID(id, accountID int64) error {
	query := `
		DELETE FROM
			transaction
		WHERE
			id = $1
			AND account_id = $2
	`
	_, err := tr.db.Exec(query, id, accountID)
	return err
}
