package repository

import (
	"user/model"

	"gorm.io/gorm"
)

type MysqlUserRepository struct {
	Conn *gorm.DB
}

func NewUserRepository(Conn *gorm.DB) model.UserRepository {
	return &MysqlUserRepository{Conn: Conn}
}
func (m *MysqlUserRepository) RegisterUser(u *model.User) (res model.User, err error) {
	result := m.Conn.Create(&u)
	return res, result.Error
}
func (m *MysqlUserRepository) LoginUser(username string, password string) (err error) {
	var dbuser model.User
	result := m.Conn.Where("username = ? AND password = ?", username, password).First(&dbuser)
	return result.Error
}
