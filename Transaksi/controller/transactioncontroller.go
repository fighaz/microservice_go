package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"transaksi/auth"
	"transaksi/model"

	"github.com/gorilla/mux"
)

type TransactionController struct {
	TransactionUsecase model.TransactionUsecase
}

func TransactionRouter(TransactionUsecase model.TransactionUsecase) {
	r := mux.NewRouter()
	http.Handle("/", r)
	r.Use(auth.IsAunthenticate)
	handler := &TransactionController{TransactionUsecase: TransactionUsecase}
	r.HandleFunc("/insert", handler.CreateTransaction).Methods("POST")
	r.HandleFunc("/transaction/{id}", handler.GetTransactionById).Methods("GET")
}
func (tc *TransactionController) GetTransactionById(w http.ResponseWriter, r *http.Request) {
	var err error
	var response model.Response

	vars := mux.Vars(r)
	getid := vars["id"]
	id, err := strconv.Atoi(getid)

	transaction, err := tc.TransactionUsecase.GetTransactionById(r.Context(), id)
	if err != nil {
		response.Message = err.Error()
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(response)
	} else {
		response.Message = "Get Data Success"
		response.Data = transaction
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		json.NewEncoder(w).Encode(response)
	}

}
func (tc *TransactionController) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var err error
	var response model.Response
	var ctx context.Context
	var transaction model.Transaction

	token := auth.GetToken(r)
	ctx = context.WithValue(r.Context(), "token", token)

	err = json.NewDecoder(r.Body).Decode(&transaction)
	if err != nil {
		response.Message = err.Error()
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(response)
	}
	transaction, err = tc.TransactionUsecase.CreateTransaction(ctx, &transaction)
	if err != nil {
		response.Message = err.Error()
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(response)
	} else {
		response.Message = "Get Data Success"
		response.Data = transaction
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		json.NewEncoder(w).Encode(response)
	}

}
