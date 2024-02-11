package repository

import "github.com/gustapinto/go_hex/internal/core/entity"

type Account interface {
	GetByID(int64) (entity.Account, error)

	GetAll() ([]entity.Account, error)

	Create(string, float64) (int64, error)

	UpdateByID(int64, string, float64) error

	DeleteByID(int64) error
}
