package model

type Product struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	Stock int    `json:"stock"`
}
type ProductUsecase interface {
	InsertProduct(p *Product) (Product, error)
	GetProductById(id int) (Product, error)
}
type ProductRepository interface {
	InsertProduct(p *Product) (Product, error)
	GetProductById(id int) (Product, error)
}
