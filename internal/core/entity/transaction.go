package entity

import "time"

type Transaction struct {
	ID        int64
	Name      string
	AccountID int64
	Value     float64
	CreatedAt time.Time
}
