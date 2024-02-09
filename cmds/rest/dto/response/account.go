package response

import (
	"time"

	"github.com/gustapinto/go_hex/internal/core/entity"
)

type Account struct {
	ID           int64     `json:"id"`
	Name         string    `json:"name"`
	InitialValue float64   `json:"initial_value"`
	CurrentValue float64   `json:"current_value"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func NewAccountFromEntity(account entity.Account) Account {
	return Account{
		ID:           account.ID,
		Name:         account.Name,
		InitialValue: account.InitialValue,
		CurrentValue: account.CurrentValue,
		CreatedAt:    account.CreatedAt,
		UpdatedAt:    account.UpdatedAt,
	}
}

func NewAccountSliceFromEntity(accounts []entity.Account) []Account {
	var acs []Account
	for _, account := range accounts {
		acs = append(acs, NewAccountFromEntity(account))
	}

	return acs
}
