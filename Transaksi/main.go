package main

import (
	"fmt"
	"net/http"
	"transaksi/config"
	"transaksi/controller"
	"transaksi/repository"
	"transaksi/usecase"
)

func main() {
	db := config.Connect()

	transactionrepo := repository.NewTransactionRepository(db)
	transactionusecase := usecase.NewTransactionUsecase(transactionrepo)
	controller.TransactionRouter(transactionusecase)
	fmt.Println("server started at localhost:8000")
	http.ListenAndServe(":8000", nil)
}
