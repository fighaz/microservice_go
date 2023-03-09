package usecase

import (
	"user/auth"
	"user/model"
)

var secretkey string = "secretkeyjwt"

type UserUsecase struct {
	UserRepo model.UserRepository
}

func NewUserUsecase(UserRepo model.UserRepository) model.UserUsecase {
	return &UserUsecase{UserRepo: UserRepo}
}
func (uc *UserUsecase) RegisterUser(u *model.User) (res model.User, err error) {
	// res, err = uc.UserRepo.LoginUser(u)
	// u.Password, err = auth.GeneratehashPassword(u.Password)
	res, err = uc.UserRepo.RegisterUser(u)
	return res, err
}
func (uc *UserUsecase) LoginUser(username string, password string) (token string, err error) {
	// password, err = auth.GeneratehashPassword(password)
	// if err != nil {
	// 	return res, err
	// }
	err = uc.UserRepo.LoginUser(username, password)
	if err != nil {
		return token, err
	}
	token, err = auth.GenerateJWT(username)
	if err != nil {
		return token, err
	}
	return token, nil

}
