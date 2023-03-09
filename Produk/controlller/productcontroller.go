package controlller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"produk/auth"
	"produk/model"
	"strconv"

	"github.com/gorilla/mux"
)

type ProductController struct {
	ProductUsecase model.ProductUsecase
}

func ProductRouter(ProductUsecase model.ProductUsecase) {

	r := mux.NewRouter()
	r.Use(auth.IsAunthenticate)
	http.Handle("/", r)
	handler := &ProductController{ProductUsecase: ProductUsecase}
	r.HandleFunc("/insert", handler.InsertProduct).Methods("POST")
	r.HandleFunc("/product/{id}", handler.GetProductById).Methods("GET")
}

func (pu *ProductController) InsertProduct(w http.ResponseWriter, r *http.Request) {
	var response model.ResponseProduct
	var p model.Product
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		response.Message = err.Error()
	}
	_, err = pu.ProductUsecase.InsertProduct(&p)
	if err != nil {
		response.Message = err.Error()
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}
	fmt.Println("Insert Succes")

	response.Message = "Insert Succes"
	response.Data = p
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
}
func (pu *ProductController) GetProductById(w http.ResponseWriter, r *http.Request) {
	var response model.ResponseProduct
	var res model.Product
	vars := mux.Vars(r)

	getid := vars["id"]
	id, err := strconv.Atoi(getid)

	res, err = pu.ProductUsecase.GetProductById(id)
	if err != nil {
		response.Message = err.Error()
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(response)

	} else {
		fmt.Println("succes")
		response.Message = "Get Data Success"
		response.Data = res
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		json.NewEncoder(w).Encode(response)
	}

}
