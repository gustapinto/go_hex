package account

import (
	"database/sql"
	"errors"
	"fmt"
)

type Service struct {
	Repository Repository
}

var (
	ErrNotFoundByID = "error.account.not.found.id: %d"
	ErrInternal     = "error.account.internal: %s"
)

func (s Service) GetByID(id int64) (account Account, err error) {
	account, err = s.Repository.GetByID(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = fmt.Errorf(ErrNotFoundByID, id)
			return
		}

		err = fmt.Errorf(ErrInternal, err.Error())
		return
	}
	return
}
