package model

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

var Mykey = []byte("SecretEncoder")

type Claims struct {
	Username string
	Role     bool
	Fname    string
	jwt.StandardClaims
}

func IsValidCoockie(w http.ResponseWriter, r *http.Request) bool {
	fmt.Println("inside Coockie")
	cookie, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return false
		}
		w.WriteHeader(http.StatusBadRequest)
		return false
	}
	tokenStr := cookie.Value
	claims := &Claims{}
	fmt.Println("token is :", tokenStr)
	fmt.Println("Claims data:", claims)

	tkn, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {

		return Mykey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return false
		}
		w.WriteHeader(http.StatusBadRequest)
		return false
	}
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return false
	}
	return true
}


