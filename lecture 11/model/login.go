package model

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var ActiveUser User

func LoginPage(w http.ResponseWriter, r *http.Request) {
	var IndexOfUser = -1
	var userforlogin User
	var password string
	err := json.NewDecoder(r.Body).Decode(&userforlogin)
	password = userforlogin.Password
	fmt.Println("Inside Login")
	fmt.Println("data frontend uname", userforlogin.UniqueID)
	fmt.Println("data user uname", AllUser[0].UniqueID)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	for i := 0; i < len(AllUser); i++ {

		if AllUser[i].UniqueID == userforlogin.UniqueID {
			IndexOfUser = i

			break
		}
	}
	if IndexOfUser == -1 {
		fmt.Println(userforlogin)
		if password != userforlogin.Password {
			w.WriteHeader(http.StatusUnauthorized)
			return
		} else {
			fmt.Fprintf(w, "%s", "Wrong Password")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
	}

	expirationTime := time.Now().Add(time.Minute * 20)
	claims := &Claims{
		Username: AllUser[IndexOfUser].UniqueID,
		Role:     AllUser[IndexOfUser].IsAdmin,
		Fname:    AllUser[IndexOfUser].Fname,

		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(Mykey)
	if err != nil {

		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Println(tokenString)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
	json.NewEncoder(w).Encode(AllUser[IndexOfUser].IsAdmin)
	ActiveUser = AllUser[IndexOfUser]
	fmt.Println(" Index of user in login:", IndexOfUser)

}
