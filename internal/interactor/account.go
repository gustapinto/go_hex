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

	ErrTransactionNotFound               = "error.transaction.not.found"
	ErrTransactionInternal               = "error.transaction.internal: %s"
	ErrTransactionValidationInvalidName  = "error.account.validation.name: %s"
	ErrTransactionValidationInvalidValue = "error.account.validation.initial.value: %f"
)

type Account struct {
	repository            repository.Account
	transactionRepository repository.Transaction
}

func NewAccount(repository repository.Account, transactionRepository repository.Transaction) Account {
	return Account{
		repository:            repository,
		transactionRepository: transactionRepository,
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

func (Account) validateTransactionName(name string) error {
	if strings.TrimSpace(name) == "" {
		return fmt.Errorf(ErrTransactionValidationInvalidName, name)
	}

	return nil
}

func (Account) validateTransactionValue(value float64) error {
	if value < 0.0 {
		return fmt.Errorf(ErrTransactionValidationInvalidValue, value)
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

func (in Account) GetTransactionByIDAndAccountID(id, accountID int64) (transaction entity.Transaction, err error) {
	transaction, err = in.transactionRepository.GetByIDAndAccountID(id, accountID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return transaction, fmt.Errorf(ErrTransactionNotFound)
		}

		return transaction, fmt.Errorf(ErrTransactionInternal, err.Error())
	}
	return
}

func (in Account) GetTransactionsByAccountID(id int64) (transactions []entity.Transaction, err error) {
	transactions, err = in.transactionRepository.GetAllByAccountID(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf(ErrTransactionNotFound)
		}

		return nil, fmt.Errorf(ErrTransactionInternal, err.Error())
	}

	return transactions, nil
}

func (in Account) CreateTransaction(accountID int64, name string, value float64) (id int64, err error) {
	if err := in.validateTransactionName(name); err != nil {
		return 0, err
	}

	if err := in.validateTransactionValue(value); err != nil {
		return 0, err
	}

	id, err = in.transactionRepository.CreateByAccountID(accountID, value, name)
	if err != nil {
		return 0, fmt.Errorf(ErrTransactionInternal, err.Error())
	}

	account, err := in.GetByID(accountID)
	if err != nil {
		return 0, err
	}

	currentAccountValue := account.CurrentValue + value
	if err := in.UpdateByID(accountID, account.Name, currentAccountValue); err != nil {
		return 0, err
	}

	return id, nil
}

func (in Account) DeleteTransaction(id, accountID int64) error {
	transaction, err := in.GetTransactionByIDAndAccountID(id, accountID)
	if err != nil {
		return err
	}

	account, err := in.GetByID(accountID)
	if err != nil {
		return err
	}

	currentAccountValue := account.CurrentValue + (transaction.Value * -1)
	if err := in.UpdateByID(accountID, account.Name, currentAccountValue); err != nil {
		return err
	}

	return nil
}
