package server

import (
	"context"

	"github.com/gustapinto/go_hex/cmds/grpc/gen"
	"github.com/gustapinto/go_hex/internal/core/entity"
	"github.com/gustapinto/go_hex/internal/interactor"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var _ gen.TransactionServiceServer = (*Transaction)(nil)

type Transaction struct {
	interactor interactor.Account
}

func NewTransaction(interactor interactor.Account) Transaction {
	return Transaction{
		interactor: interactor,
	}
}

func (t *Transaction) Create(_ context.Context, req *gen.CreateTransactionRequest) (res *gen.CreatedResponse, err error) {
	if req == nil {
		return nil, ErrEmptyRequest()
	}

	id, err := t.interactor.CreateTransaction(req.AccountId, req.Name, req.Value)
	if err != nil {
		return nil, ErrInternal(err)
	}

	res = &gen.CreatedResponse{
		Id: id,
	}
	return
}

func (t *Transaction) DeleteByIDAndAccountID(_ context.Context, req *gen.DeleteTransactionByIDAndAccountIDRequest) (_ *gen.Empty, err error) {
	if req == nil {
		return nil, ErrEmptyRequest()
	}

	err = t.interactor.DeleteTransaction(req.Id, req.AccountId)
	if err != nil {
		return nil, ErrInternal(err)
	}

	return nil, nil
}

func (*Transaction) convertTransactionToTransactionResponse(transaction entity.Transaction) *gen.TransactionResponse {
	if transaction.IsZero() {
		return nil
	}

	return &gen.TransactionResponse{
		Id:        transaction.ID,
		AccountId: transaction.AccountID,
		Name:      transaction.Name,
		Value:     transaction.Value,
		CreatedAt: timestamppb.New(transaction.CreatedAt),
	}
}

func (t *Transaction) GetByAccountID(_ context.Context, req *gen.GetTransactionByAccountIDRequest) (res *gen.RepeatedTransactionResponse, err error) {
	if req == nil {
		return nil, ErrEmptyRequest()
	}

	transactions, err := t.interactor.GetTransactionsByAccountID(req.AccountId)
	if err != nil {
		return nil, ErrInternal(err)
	}

	res = &gen.RepeatedTransactionResponse{
		Transactions: make([]*gen.TransactionResponse, 0),
	}

	for _, transaction := range transactions {
		res.Transactions = append(
			res.Transactions,
			t.convertTransactionToTransactionResponse(transaction))
	}
	return
}

func (t *Transaction) GetByIdAndAccountId(_ context.Context, req *gen.GetTransactionByIDAndAccountIDRequest) (res *gen.TransactionResponse, err error) {
	if req == nil {
		return nil, ErrEmptyRequest()
	}

	transaction, err := t.interactor.GetTransactionByIDAndAccountID(req.Id, req.AccountId)
	if err != nil {
		return nil, ErrInternal(err)
	}

	res = t.convertTransactionToTransactionResponse(transaction)
	return
}
