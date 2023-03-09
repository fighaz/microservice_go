package usecase

import (
	"context"
	"fmt"
	"transaksi/model"
)

type TransactionUsecase struct {
	TransactionRepo model.TransactionRepository
}

func NewTransactionUsecase(TransactionRepo model.TransactionRepository) model.TransactionUsecase {
	return &TransactionUsecase{TransactionRepo: TransactionRepo}
}
func (tu *TransactionUsecase) CreateTransaction(ctx context.Context, t *model.Transaction) (res model.Transaction, err error) {
	res, err = tu.TransactionRepo.CreateTransaction(ctx, t)
	_, err = tu.TransactionRepo.GetProductById(ctx, res.Id_Product)
	if err != nil {
		return res, err
	}
	return res, nil
}
func (tu *TransactionUsecase) GetTransactionById(ctx context.Context, id int) (res model.ResultTransaction, err error) {

	transaction, err := tu.TransactionRepo.GetTransactionById(ctx, id)
	if err != nil {
		return res, err
	}
	idprod := transaction.Id_Product
	product, err := tu.TransactionRepo.GetProductById(ctx, idprod)
	if err != nil {
		return res, err
	}
	res.Id = transaction.Id
	res.Data = product.Data

	fmt.Println(transaction.Id)
	return res, nil
}
