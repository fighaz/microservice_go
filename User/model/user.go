package model

type User struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}
type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserRepository interface {
	RegisterUser(u *User) (User, error)
	LoginUser(username string, password string) (err error)
}
type UserUsecase interface {
	RegisterUser(u *User) (res User, err error)
	LoginUser(username string, password string) (token string, err error)
}
