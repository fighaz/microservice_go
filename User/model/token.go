package model

type Token struct {
	Username string `json:"username"`
	JWTToken string `json:"token"`
}
