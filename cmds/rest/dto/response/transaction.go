package response

import (
	"time"

	"github.com/gustapinto/go_hex/internal/core/entity"
)

type Transaction struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	AccountID int64     `json:"account_id"`
	Value     float64   `json:"value"`
	CreatedAt time.Time `json:"created_at"`
}

func NewTransactionFromEntity(transaction entity.Transaction) Transaction {
	return Transaction{
		ID:        transaction.ID,
		Name:      transaction.Name,
		AccountID: transaction.AccountID,
		Value:     transaction.Value,
		CreatedAt: transaction.CreatedAt,
	}
}

func NewTransactionSliceFromEntity(transactions []entity.Transaction) []Transaction {
	trs := make([]Transaction, len(transactions))

	for i, entity := range transactions {
		trs[i] = NewTransactionFromEntity(entity)
	}

	return trs
}
