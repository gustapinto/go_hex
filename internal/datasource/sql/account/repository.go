package account

import (
	"database/sql"

	"github.com/gustapinto/go_hex/internal/core/account"
)

type RepositoryImpl struct {
	DB *sql.DB
}

var _ account.Repository = (*RepositoryImpl)(nil) // Validate interface compliance on compile time

func (ri RepositoryImpl) GetByID(id int64) (account account.Account, err error) {
	// TODO
	return
}

func (ri RepositoryImpl) GetAll() (accounts []account.Account, err error) {
	// TODO
	return
}

func (ri RepositoryImpl) Count() (total int64, err error) {
	// TODO
	return
}

func (ri RepositoryImpl) Create(name string, initialValue float64) (id int64, err error) {
	// TODO
	return
}

func (ri RepositoryImpl) UpdateByID(id int64, name string, currentValue float64) error {
	// TODO
	return nil
}

func (ri RepositoryImpl) SumToCurrentValue(id int64, value float64) error {
	// TODO
	return nil
}

func (ri RepositoryImpl) SubtractFromCurrentValue(id int64, value float64) error {
	// TODO
	return nil
}

func (ri RepositoryImpl) DeleteByID(id int64) error {
	// TODO
	return nil
}
