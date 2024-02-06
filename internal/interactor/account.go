package interactor

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/gustapinto/go_hex/internal/core/entity"
	"github.com/gustapinto/go_hex/internal/core/repository"
)

var (
	ErrNotFound                      = "error.account.not.found"
	ErrInternal                      = "error.account.internal: %s"
	ErrValidationInvalidName         = "error.account.validation.name: %s"
	ErrValidationInvalidInitialValue = "error.account.validation.initial.value: %f"
)

type Account struct {
	repository repository.Account
}

func NewAccount(repository repository.Account) Account {
	return Account{
		repository: repository,
	}
}

func (Account) validateName(name string) error {
	if strings.TrimSpace(name) == "" {
		return fmt.Errorf(ErrValidationInvalidName, name)
	}

	return nil
}

func (Account) validateInitialValue(initialValue float64) error {
	if initialValue < 0.0 {
		return fmt.Errorf(ErrValidationInvalidInitialValue, initialValue)
	}

	return nil
}

func (in Account) GetByID(id int64) (account entity.Account, err error) {
	account, err = in.repository.GetByID(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return account, fmt.Errorf(ErrNotFound, id)
		}

		return account, fmt.Errorf(ErrInternal, err.Error())
	}
	return
}

func (in Account) GetAll() (accounts []entity.Account, err error) {
	accounts, err = in.repository.GetAll()
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return accounts, fmt.Errorf(ErrNotFound)
		}

		return accounts, fmt.Errorf(ErrInternal, err.Error())
	}
	return
}

func (in Account) Create(name string, initialValue float64) (id int64, err error) {
	if err := in.validateName(name); err != nil {
		return 0, err
	}

	if err := in.validateInitialValue(initialValue); err != nil {
		return 0, err
	}

	id, err = in.repository.Create(name, initialValue)
	if err != nil {
		return 0, fmt.Errorf(ErrInternal, err.Error())
	}

	return id, nil
}

func (in Account) UpdateByID(id int64, name string, currentValue float64) error {
	if err := in.validateName(name); err != nil {
		return err
	}

	if err := in.repository.UpdateByID(id, name, currentValue); err != nil {
		return fmt.Errorf(ErrInternal, err.Error())
	}

	return nil
}

func (in Account) DeleteByID(id int64) error {
	if err := in.repository.DeleteByID(id); err != nil {
		return fmt.Errorf(ErrInternal, err.Error())
	}

	return nil
}
