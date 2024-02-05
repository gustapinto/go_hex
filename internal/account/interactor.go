package account

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
)

var (
	ErrNotFound                      = "error.account.not.found"
	ErrInternal                      = "error.account.internal: %s"
	ErrValidationInvalidName         = "error.account.validation.name: %s"
	ErrValidationInvalidInitialValue = "error.account.validation.initial.value: %f"
)

type Interactor struct {
	repository Repository
}

func NewInteractor(repository Repository) Interactor {
	return Interactor{
		repository: repository,
	}
}

func (in Interactor) GetByID(id int64) (account Account, err error) {
	account, err = in.repository.GetByID(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return account, fmt.Errorf(ErrNotFound, id)
		}

		return account, fmt.Errorf(ErrInternal, err.Error())
	}
	return
}

func (in Interactor) GetAll() (accounts []Account, total int64, err error) {
	accounts, err = in.repository.GetAll()
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return accounts, total, fmt.Errorf(ErrNotFound)
		}

		return accounts, total, fmt.Errorf(ErrInternal, err.Error())
	}

	total, err = in.repository.Count()
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return accounts, total, fmt.Errorf(ErrNotFound)
		}

		return accounts, total, fmt.Errorf(ErrInternal, err.Error())
	}
	return
}

func (in Interactor) Create(name string, initialValue float64) (id int64, err error) {
	if strings.TrimSpace(name) == "" {
		return 0, fmt.Errorf(ErrValidationInvalidName, name)
	}

	if initialValue < 0.0 {
		return 0, fmt.Errorf(ErrValidationInvalidInitialValue, initialValue)
	}

	id, err = in.repository.Create(name, initialValue)
	if err != nil {
		return 0, fmt.Errorf(ErrInternal, err.Error())
	}

	return id, nil
}
