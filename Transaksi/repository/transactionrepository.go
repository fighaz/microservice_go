package repository

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"transaksi/model"

	"gorm.io/gorm"
)

type MysqlTransactionRepository struct {
	Conn *gorm.DB
}

func NewTransactionRepository(Conn *gorm.DB) model.TransactionRepository {
	return &MysqlTransactionRepository{Conn: Conn}
}
func (m *MysqlTransactionRepository) CreateTransaction(ctx context.Context, t *model.Transaction) (res model.Transaction, err error) {
	result := m.Conn.Create(&t)
	return res, result.Error
}
func (m *MysqlTransactionRepository) GetTransactionById(ctx context.Context, id int) (res model.Transaction, err error) {
	result := m.Conn.First(&res, id)
	return res, result.Error
}
func (m *MysqlTransactionRepository) GetProductById(ctx context.Context, id int) (res model.ResponseProduct, err error) {
	var client = &http.Client{}
	var product model.ResponseProduct
	param := strconv.Itoa(id)
	token := ctx.Value("token").(string)
	log.Println(token)
	request, err := http.NewRequest("GET", "http://localhost:9000/product/"+param, nil)
	request.Header.Set("Authorization", "Bearer "+token)
	if err != nil {
		return res, err
	}
	data, err := client.Do(request)
	if err != nil {

		return res, err
	}
	defer data.Body.Close()
	err = json.NewDecoder(data.Body).Decode(&product)
	if err != nil {
		return res, err
	}
	res = product
	return res, nil
}
