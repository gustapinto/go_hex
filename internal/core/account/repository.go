package account

type Repository interface {
	GetByID(int64) (Account, error)

	GetAll() ([]Account, int64, error)

	Create(int64, string, float64) (int64, error)

	UpdateByID(int64, string, float64) error

	DeleteByID(int64) error
}
