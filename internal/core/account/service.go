package account

import (
	"database/sql"
	"errors"
	"fmt"
)

type Service struct {
	Repository Repository
}

var (
	ErrNotFound = "error.account.not.found"
	ErrInternal = "error.account.internal: %s"
)

func (s Service) GetByID(id int64) (account Account, err error) {
	account, err = s.Repository.GetByID(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return account, fmt.Errorf(ErrNotFound, id)
		}

		return account, fmt.Errorf(ErrInternal, err.Error())
	}
	return
}

func (s Service) GetAll() (accounts []Account, total int64, err error) {
	accounts, err = s.Repository.GetAll()
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return accounts, total, fmt.Errorf(ErrNotFound)
		}

		return accounts, total, fmt.Errorf(ErrInternal, err.Error())
	}

	total, err = s.Repository.Count()
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return accounts, total, fmt.Errorf(ErrNotFound)
		}

		return accounts, total, fmt.Errorf(ErrInternal, err.Error())
	}
	return
}
