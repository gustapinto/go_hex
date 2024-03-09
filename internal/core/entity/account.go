package entity

import (
	"time"
)

type Account struct {
	ID           int64
	Name         string
	InitialValue float64
	CurrentValue float64
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (a Account) IsZero() bool {
	return a.ID == 0 &&
		a.Name == "" &&
		a.InitialValue == 0.0 &&
		a.CurrentValue == 0.0 &&
		a.CreatedAt.IsZero() &&
		a.UpdatedAt.IsZero()
}
