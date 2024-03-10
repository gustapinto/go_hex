package entity

import "time"

type Transaction struct {
	ID        int64
	Name      string
	AccountID int64
	Value     float64
	CreatedAt time.Time
}

func (t Transaction) IsZero() bool {
	return t.ID == 0 &&
		t.Name == "" &&
		t.AccountID == 0 &&
		t.Value == 0.0 &&
		t.CreatedAt.IsZero()
}
