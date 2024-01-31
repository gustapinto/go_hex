package account

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
	DeletedAt    time.Time
}
