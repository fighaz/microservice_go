package usecase

import "produk/model"

type ProductUsecase struct {
	ProductRepo model.ProductRepository
}

func NewProductUsecase(ProductRepo model.ProductRepository) model.ProductUsecase {
	return &ProductUsecase{ProductRepo: ProductRepo}
}
func (pu *ProductUsecase) InsertProduct(p *model.Product) (res model.Product, err error) {
	res, err = pu.ProductRepo.InsertProduct(p)
	return res, err
}
func (pu *ProductUsecase) GetProductById(id int) (res model.Product, err error) {
	res, err = pu.ProductRepo.GetProductById(id)
	return res, err
}
