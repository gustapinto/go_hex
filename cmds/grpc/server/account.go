package server

import (
	"context"

	"github.com/gustapinto/go_hex/cmds/grpc/gen"
	"github.com/gustapinto/go_hex/internal/core/entity"
	"github.com/gustapinto/go_hex/internal/interactor"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var _ gen.AccountServiceServer = (*Account)(nil)

type Account struct {
	interactor interactor.Account
}

func NewAccount(interactor interactor.Account) Account {
	return Account{
		interactor: interactor,
	}
}

func (a *Account) Create(_ context.Context, req *gen.CreateAccountRequest) (res *gen.CreatedResponse, err error) {
	if req == nil {
		return nil, ErrEmptyRequest()
	}

	id, err := a.interactor.Create(req.Name, req.InitalValue)
	if err != nil {
		return nil, ErrInternal(err)
	}

	res = &gen.CreatedResponse{
		Id: id,
	}
	return
}

func (a *Account) UpdateByID(_ context.Context, req *gen.UpdateAccountByIDRequest) (res *gen.Empty, err error) {
	if req == nil {
		return nil, ErrEmptyRequest()
	}

	err = a.interactor.UpdateByID(req.Id, req.Name, req.CurrentValue)
	return
}

func (a *Account) DeleteByID(_ context.Context, req *gen.DeleteAccountByIDRequest) (res *gen.Empty, err error) {
	if req == nil {
		return nil, ErrEmptyRequest()
	}

	err = a.interactor.DeleteByID(req.Id)
	if err != nil {
		return nil, ErrInternal(err)
	}

	return nil, nil
}

func (*Account) convertAccountToAccountResponse(account entity.Account) *gen.AccountResponse {
	if account.IsZero() {
		return nil
	}

	return &gen.AccountResponse{
		Id:           account.ID,
		Name:         account.Name,
		InitialValue: account.InitialValue,
		CurrentValue: account.CurrentValue,
		CreatedAt:    timestamppb.New(account.CreatedAt),
		UpdatedAt:    timestamppb.New(account.UpdatedAt),
	}
}

func (a *Account) GetAll(_ context.Context, _ *gen.Empty) (res *gen.RepeatedAccountResponse, err error) {
	accounts, err := a.interactor.GetAll()
	if err != nil {
		return nil, ErrInternal(err)
	}

	res = &gen.RepeatedAccountResponse{
		Accounts: make([]*gen.AccountResponse, len(accounts)),
	}

	for _, account := range accounts {
		res.Accounts = append(
			res.Accounts,
			a.convertAccountToAccountResponse(account))
	}
	return
}

func (a *Account) GetByID(_ context.Context, req *gen.GetAccountByIDRequest) (res *gen.AccountResponse, err error) {
	if req == nil {
		return nil, ErrEmptyRequest()
	}

	account, err := a.interactor.GetByID(req.Id)
	if err != nil {
		return nil, ErrInternal(err)
	}

	res = a.convertAccountToAccountResponse(account)
	return
}
