package account

import (
	"database/sql"
	"errors"
	"fmt"
)

var (
	ErrNotFound = "error.account.not.found"
	ErrInternal = "error.account.internal: %s"
)

type Interactor struct {
	repository Repository
}

func NewInteractor(repository Repository) Interactor {
	return Interactor{
		repository: repository,
	}
}

func (s Interactor) GetByID(id int64) (account Account, err error) {
	account, err = s.repository.GetByID(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return account, fmt.Errorf(ErrNotFound, id)
		}

		return account, fmt.Errorf(ErrInternal, err.Error())
	}
	return
}

func (s Interactor) GetAll() (accounts []Account, total int64, err error) {
	accounts, err = s.repository.GetAll()
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return accounts, total, fmt.Errorf(ErrNotFound)
		}

		return accounts, total, fmt.Errorf(ErrInternal, err.Error())
	}

	total, err = s.repository.Count()
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return accounts, total, fmt.Errorf(ErrNotFound)
		}

		return accounts, total, fmt.Errorf(ErrInternal, err.Error())
	}
	return
}
