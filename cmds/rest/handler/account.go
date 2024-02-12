package handler

import (
	"net/http"

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
	id, err := httputil.PathValueInt64(w, r, "accountID")
	if err != nil {
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

	id, err := httputil.PathValueInt64(w, r, "accountID")
	if err != nil {
		return
	}

	if err := ac.interactor.UpdateByID(id, req.Name, req.CurrentValue); err != nil {
		httputil.WriteJson(w, r, http.StatusBadRequest, httputil.NewErrorResponse(err))
		return
	}

	httputil.WriteStatusCode(w, r, http.StatusNoContent)
}

func (ac Account) DeletebyID(w http.ResponseWriter, r *http.Request) {
	id, err := httputil.PathValueInt64(w, r, "accountID")
	if err != nil {
		return
	}

	if err := ac.interactor.DeleteByID(id); err != nil {
		httputil.WriteJson(w, r, http.StatusBadRequest, httputil.NewErrorResponse(err))
		return
	}

	httputil.WriteStatusCode(w, r, http.StatusNoContent)
}

func (ac Account) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var req request.CreateTransaction
	if err := httputil.BindJson(w, r, &req); err != nil {
		return
	}

	accountID, err := httputil.PathValueInt64(w, r, "accountID")
	if err != nil {
		return
	}

	id, err := ac.interactor.CreateTransaction(accountID, req.Name, req.Value)
	if err != nil {
		httputil.WriteJson(w, r, http.StatusBadRequest, httputil.NewErrorResponse(err))
		return
	}

	httputil.WriteJson(w, r, http.StatusCreated, httputil.NewCreatedResponse(id))
}

func (ac Account) GetTransactionsByAccountID(w http.ResponseWriter, r *http.Request) {
	accountID, err := httputil.PathValueInt64(w, r, "accountID")
	if err != nil {
		return
	}

	transactions, err := ac.interactor.GetTransactionsByAccountID(accountID)
	if err != nil {
		httputil.WriteJson(w, r, http.StatusBadRequest, httputil.NewErrorResponse(err))
		return
	}

	httputil.WriteJson(w, r, http.StatusCreated, response.NewTransactionSliceFromEntity(transactions))
}

func (ac Account) GetTransactionByIDAndAccountID(w http.ResponseWriter, r *http.Request) {
	accountID, err := httputil.PathValueInt64(w, r, "accountID")
	if err != nil {
		return
	}

	transactionID, err := httputil.PathValueInt64(w, r, "transactionID")
	if err != nil {
		return
	}

	transaction, err := ac.interactor.GetTransactionByIDAndAccountID(transactionID, accountID)
	if err != nil {
		httputil.WriteJson(w, r, http.StatusBadRequest, httputil.NewErrorResponse(err))
		return
	}

	httputil.WriteJson(w, r, http.StatusCreated, response.NewTransactionFromEntity(transaction))
}

func (ac Account) DeleteTransactionByIDAndAccountID(w http.ResponseWriter, r *http.Request) {
	accountID, err := httputil.PathValueInt64(w, r, "accountID")
	if err != nil {
		return
	}

	transactionID, err := httputil.PathValueInt64(w, r, "transactionID")
	if err != nil {
		return
	}

	if err = ac.interactor.DeleteTransaction(transactionID, accountID); err != nil {
		httputil.WriteJson(w, r, http.StatusBadRequest, httputil.NewErrorResponse(err))
		return
	}

	httputil.WriteStatusCode(w, r, http.StatusNoContent)
}
