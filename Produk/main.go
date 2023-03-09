package main

import (
	"fmt"
	"net/http"
	"produk/config"
	"produk/controlller"
	"produk/repository"
	"produk/usecase"
)

func main() {
	db := config.Connect()

	productrepo := repository.NewProductRepository(db)
	productusecase := usecase.NewProductUsecase(productrepo)
	controlller.ProductRouter(productusecase)

	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}
