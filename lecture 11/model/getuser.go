package model

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

func GetUser(w http.ResponseWriter, r *http.Request) string {
	//refreshToken(w, r)
	cookie, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return ""
		}
		w.WriteHeader(http.StatusBadRequest)
		return ""
	}
	tokenStr := cookie.Value
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
		return t, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return ""
		}
		w.WriteHeader(http.StatusBadRequest)
		return ""
	}
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return ""
	}

	fmt.Fprintf(w, "Hello, %s", claims.Username)
	return claims.Username
}
