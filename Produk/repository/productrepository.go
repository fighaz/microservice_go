package repository

import (
	"produk/model"

	"gorm.io/gorm"
)

type MysqlProductRepository struct {
	Conn *gorm.DB
}

func NewProductRepository(Conn *gorm.DB) model.ProductRepository {
	return &MysqlProductRepository{Conn}
}
func (m *MysqlProductRepository) InsertProduct(p *model.Product) (res model.Product, err error) {
	result := m.Conn.Create(&p)
	return res, result.Error
}
func (m *MysqlProductRepository) GetProductById(id int) (res model.Product, err error) {
	result := m.Conn.First(&res, id)
	return res, result.Error
}
