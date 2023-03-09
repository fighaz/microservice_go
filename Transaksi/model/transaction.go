package model

import "context"

type Transaction struct {
	Id         int `json:"id"`
	Id_Product int `json:"id_product"`
}
type ResultTransaction struct {
	Id   int
	Data Product
}
type TransactionUsecase interface {
	CreateTransaction(ctx context.Context, t *Transaction) (Transaction, error)
	GetTransactionById(ctx context.Context, id int) (ResultTransaction, error)
}
type TransactionRepository interface {
	CreateTransaction(ctx context.Context, t *Transaction) (Transaction, error)
	GetTransactionById(ctx context.Context, id int) (Transaction, error)
	GetProductById(ctx context.Context, id int) (ResponseProduct, error)
}
