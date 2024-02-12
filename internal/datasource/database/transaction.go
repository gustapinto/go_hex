package database

import (
	"database/sql"

	"github.com/gustapinto/go_hex/internal/core/entity"
	"github.com/gustapinto/go_hex/internal/core/repository"

	_ "embed"
)

var (
	// Validate interface compliance on compile time
	_ repository.Transaction = (*Transaction)(nil)

	//go:embed query/transaction/get_by_id_and_account_id.sql
	transactionGetByIDAndAccountIDQuery string

	//go:embed query/transaction/get_all_by_account_id.sql
	transactionGetAllByAccountIDQuery string

	//go:embed query/transaction/create_by_account_id.sql
	transactionCreateByAccountIdQuery string

	//go:embed query/transaction/delete_by_id_and_account_id.sql
	transactionDeleteByIDAndAccountIDQuery string
)

type Transaction struct {
	db *sql.DB
}

func NewTransaction(db *sql.DB) Transaction {
	return Transaction{
		db: db,
	}
}

func (tr Transaction) GetByIDAndAccountID(id, accountID int64) (transaction entity.Transaction, err error) {
	row := tr.db.QueryRow(transactionGetByIDAndAccountIDQuery, id, accountID)
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
	rows, err := tr.db.Query(transactionGetAllByAccountIDQuery, accountID)
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
	row := tr.db.QueryRow(transactionCreateByAccountIdQuery, name, accountID, value)
	if row.Err() != nil {
		return 0, row.Err()
	}

	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return
}

func (tr Transaction) DeleteByIDAndAccountID(id, accountID int64) error {
	_, err := tr.db.Exec(transactionDeleteByIDAndAccountIDQuery, id, accountID)
	return err
}
