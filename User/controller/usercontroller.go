package controller

import (
	"encoding/json"
	"net/http"
	"user/model"

	"github.com/gorilla/mux"
)

type UserController struct {
	UserUsecase model.UserUsecase
}

func UserRouter(UserUsecase model.UserUsecase) {
	r := mux.NewRouter()
	http.Handle("/", r)
	handler := &UserController{UserUsecase: UserUsecase}
	r.HandleFunc("/register", handler.RegisterUser).Methods("POST")
	r.HandleFunc("/login", handler.LoginUser).Methods("POST")
}

func (uc *UserController) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var err error
	var response model.Response
	var user model.User
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		response.Message = err.Error()
	}
	_, err = uc.UserUsecase.RegisterUser(&user)
	if err != nil {

		response.Message = err.Error()
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		json.NewEncoder(w).Encode(response)
		return
	} else {
		response.Message = "Account Succes Created"
		response.Data = user
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		json.NewEncoder(w).Encode(response)
	}
}
func (uc *UserController) LoginUser(w http.ResponseWriter, r *http.Request) {
	var err error
	var response model.Response
	var login model.Login

	var token model.Token
	err = json.NewDecoder(r.Body).Decode(&login)
	if err != nil {
		response.Message = err.Error()
		return
	}
	jwt, err := uc.UserUsecase.LoginUser(login.Username, login.Password)
	if err != nil {
		response.Message = err.Error()
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(response)
		return
	}

	token.Username = login.Username
	token.JWTToken = jwt
	response.Message = "Login Success"
	response.Data = token
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(response)
}
