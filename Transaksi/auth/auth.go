package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"transaksi/model"

	"github.com/golang-jwt/jwt"
)

var secretkey string = "secretkeyjwt"

func GetToken(r *http.Request) string {
	bearerToken := r.Header.Get("Authorization")
	splitToken := strings.Split(bearerToken, " ")
	if len(splitToken) != 2 || strings.ToLower(splitToken[0]) != "bearer" {
		return ""
	}

	return splitToken[1]
}
func IsAunthenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		var response model.Response
		jwttoken := GetToken(r)
		if jwttoken == "" {
			response.Message = "No Token Found"
			json.NewEncoder(w).Encode(response)
			return
		}
		ctx := context.WithValue(r.Context(), "token", jwttoken)
		var mySigningKey = []byte(secretkey)

		token, err := jwt.Parse(jwttoken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("There was an error in parsing")
			}
			return mySigningKey, nil
		})

		if err != nil {
			response.Message = err.Error()
			json.NewEncoder(w).Encode(response)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			next.ServeHTTP(w, r.WithContext(ctx))
			json.NewEncoder(w).Encode(claims)
		} else {
			response.Message = err.Error()
			json.NewEncoder(w).Encode(response)
			return
		}

	})
}
