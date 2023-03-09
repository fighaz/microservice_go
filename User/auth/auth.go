package auth

import (
	"time"
	"user/model"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

var secretkey string = "secretkeyjwt"

func GenerateJWT(username string) (string, error) {
	var response model.Response
	var mySigningKey = []byte(secretkey)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		response.Message = err.Error()
		return "", err
	}
	return tokenString, nil
}
func GeneratehashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
