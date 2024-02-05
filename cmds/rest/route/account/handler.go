package account

import (
	"net/http"
	"strings"

	"github.com/gustapinto/go_hex/internal/account"
	"github.com/gustapinto/go_hex/pkg/httputil"
)

type Account struct {
	interactor account.Interactor
}

func NewAccount(interactor account.Interactor) Account {
	return Account{
		interactor: interactor,
	}
}

func (ac Account) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var req CreateAccountRequest

	if err := httputil.BindJson(w, r, &req); err != nil {
		return
	}

	id, err := ac.interactor.Create(req.Name, req.InitialValue)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			httputil.WriteJson(w, r, http.StatusBadRequest, httputil.ErrorResponse{
				Err: err.Error(),
			})
			return
		}

		httputil.WriteJson(w, r, http.StatusInternalServerError, httputil.ErrorResponse{
			Err: err.Error(),
		})
		return
	}

	httputil.WriteJson(w, r, http.StatusCreated, httputil.CreatedResponse{
		ID: id,
	})
}
