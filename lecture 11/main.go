package main

import (
	"fmt"
	"log"
	"model/control"
	"model/model"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	model.CreateUser("DN", "Darshan", "Naik", "Test", true)
	model.CreateContacts("DN", "rahul", "Nair", "Home", 9876554321)

	handleRequests()

}
func handleRequests() {
	headersOk := handlers.AllowCredentials()
	originsOk := handlers.AllowedOrigins([]string{os.Getenv("ORIGIN_ALLOWED")})
	methodsOk := handlers.AllowedMethods([]string{"GET", "DELETE", "HEAD", "POST", "PUT", "OPTIONS"})
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/loginpage", model.LoginPage).Methods("POST")
	for {
		fmt.Println("main function handle ", model.ActiveUser)
		fmt.Println("Is Admin", model.ActiveUser.IsAdmin)
		if model.ActiveUser.IsAdmin {
			router.HandleFunc("/readuser", control.ReadUsers).Methods("GET")              //read		USER
			router.HandleFunc("/createuser", control.CreateUser).Methods("POST")          //create
			router.HandleFunc("/updateuser/{unique}", control.UpdateUser).Methods("PUT")  //update
			router.HandleFunc("/deleteuser/{unique}", control.DeleteUsers).Methods("PUT") //delete

			router.HandleFunc("/readcontact", control.ReadAllContacts).Methods("GET")           //read		Contact
			router.HandleFunc("/createcontact", control.CreateContacts).Methods("POST")         //create
			router.HandleFunc("/updatecontact/{unique}", control.UpdateContacts).Methods("PUT") //update
			router.HandleFunc("/deletecontact/{unique}", control.DeleteContacts).Methods("PUT") //delete

			router.HandleFunc("/readcontactdetail", control.ReadContactDetail).Methods("GET")              //read	Contact detail
			router.HandleFunc("/createcontactdetail", control.CreateContactDetail).Methods("POST")         //create
			router.HandleFunc("/updatecontactdetail/{unique}", control.UpdateContactDetail).Methods("PUT") //update
			router.HandleFunc("/deletecontactdetail/{unique}", control.DeleteContactDetail).Methods("PUT") //delete

			log.Fatal(http.ListenAndServe(":4003", handlers.CORS(originsOk, headersOk, methodsOk)(router)))
		} else {
			router.HandleFunc("/readcontact", control.ReadAllContacts).Methods("GET")           //read		Contact
			router.HandleFunc("/createcontact", control.CreateContacts).Methods("POST")         //create
			router.HandleFunc("/updatecontact/{unique}", control.UpdateContacts).Methods("PUT") //update
			router.HandleFunc("/deletecontact/{unique}", control.DeleteContacts).Methods("PUT") //delete

			router.HandleFunc("/readcontactdetail", control.ReadContactDetail).Methods("GET")              //read	Contact detail
			router.HandleFunc("/createcontactdetail", control.CreateContactDetail).Methods("POST")         //create
			router.HandleFunc("/updatecontactdetail/{unique}", control.UpdateContactDetail).Methods("PUT") //update
			router.HandleFunc("/deletecontactdetail/{unique}", control.DeleteContactDetail).Methods("PUT") //delete
			log.Fatal(http.ListenAndServe(":4003", handlers.CORS(originsOk, headersOk, methodsOk)(router)))
		}
	}

}

// func homePage(w http.ResponseWriter, r *http.Request) {
// 	//refreshToken(w, r)
// 	cookie, err := r.Cookie("token")
// 	if err != nil {
// 		if err == http.ErrNoCookie {
// 			w.WriteHeader(http.StatusUnauthorized)
// 			return
// 		}
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}
// 	tokenStr := cookie.Value
// 	claims := &Claims{}
// 	tkn, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
// 		return jwtKey, nil
// 	})
// 	if err != nil {
// 		if err == jwt.ErrSignatureInvalid {
// 			w.WriteHeader(http.StatusUnauthorized)
// 			return
// 		}
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}
// 	if !tkn.Valid {
// 		w.WriteHeader(http.StatusUnauthorized)
// 		return
// 	}

// 	fmt.Fprintf(w, "Hello, %s", claims.Username)
// }
