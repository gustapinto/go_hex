package account

type Repository interface {
	GetByID(int64) (Account, error)

	GetAll() ([]Account, error)

	Count() (int64, error)

	Create(string, float64) (int64, error)

	UpdateByID(int64, string, float64) error

	SumToCurrentValue(int64, float64) error

	SubtractFromCurrentValue(int64, float64) error

	DeleteByID(int64) error
}
