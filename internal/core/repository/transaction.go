package repository

import "github.com/gustapinto/go_hex/internal/core/entity"

type Transaction interface {
	GetByIDAndAccountID(int64, int64) (entity.Transaction, error)

	GetAllByAccountID(int64) ([]entity.Transaction, error)

	CreateByAccountID(int64, float64, string) (int64, error)

	DeleteByIDAndAccountID(int64, int64) error
}
