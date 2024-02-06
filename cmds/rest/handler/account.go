package handler

import (
	"net/http"
	"strconv"

	"github.com/gustapinto/go_hex/cmds/rest/dto/request"
	"github.com/gustapinto/go_hex/cmds/rest/dto/response"
	"github.com/gustapinto/go_hex/internal/interactor"
	"github.com/gustapinto/go_hex/pkg/httputil"
)

type Account struct {
	interactor interactor.Account
}

func NewAccount(interactor interactor.Account) Account {
	return Account{
		interactor: interactor,
	}
}

func (ac Account) Create(w http.ResponseWriter, r *http.Request) {
	var req request.CreateAccount
	if err := httputil.BindJson(w, r, &req); err != nil {
		return
	}

	id, err := ac.interactor.Create(req.Name, req.InitialValue)
	if err != nil {
		httputil.WriteJson(w, r, http.StatusBadRequest, httputil.NewErrorResponse(err))
		return
	}

	httputil.WriteJson(w, r, http.StatusCreated, httputil.NewCreatedResponse(id))
}

func (ac Account) Get(w http.ResponseWriter, r *http.Request) {
	accounts, err := ac.interactor.GetAll()
	if err != nil {
		httputil.WriteJson(w, r, http.StatusBadRequest, httputil.NewErrorResponse(err))
		return
	}

	httputil.WriteJson(w, r, http.StatusOK, response.NewAccountSliceFromEntity(accounts))
}

func (ac Account) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 0)
	if err != nil {
		httputil.WriteJson(w, r, http.StatusBadRequest, httputil.NewErrorResponse(err))
		return
	}

	account, err := ac.interactor.GetByID(id)
	if err != nil {
		httputil.WriteJson(w, r, http.StatusBadRequest, httputil.NewErrorResponse(err))
		return
	}

	httputil.WriteJson(w, r, http.StatusOK, response.NewAccountFromEntity(account))
}

func (ac Account) UpdateByID(w http.ResponseWriter, r *http.Request) {
	var req request.UpdateAccount
	if err := httputil.BindJson(w, r, &req); err != nil {
		return
	}

	id, err := strconv.ParseInt(r.PathValue("id"), 10, 0)
	if err != nil {
		httputil.WriteJson(w, r, http.StatusBadRequest, httputil.NewErrorResponse(err))
		return
	}

	if err := ac.interactor.UpdateByID(id, req.Name, req.CurrentValue); err != nil {
		httputil.WriteJson(w, r, http.StatusBadRequest, httputil.NewErrorResponse(err))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (ac Account) DeletebyID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 0)
	if err != nil {
		httputil.WriteJson(w, r, http.StatusBadRequest, httputil.NewErrorResponse(err))
		return
	}

	if err := ac.interactor.DeleteByID(id); err != nil {
		httputil.WriteJson(w, r, http.StatusBadRequest, httputil.NewErrorResponse(err))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
