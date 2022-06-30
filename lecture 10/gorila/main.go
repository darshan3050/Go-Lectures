package main

import (
	"fmt"
	"log"
	"model/control"
	"model/view"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {

	view.AllUser = append(view.AllUser, *view.NewUser("darshan", "naik", "dahanu", "Darsh3050", 22, 1234567890))

	view.AllUser = append(view.AllUser, *view.NewUser("naik", "dahanu", "andheri", "Duck", 22, 9876543210))
	handleReq()
}

func handleReq() {

	headersOk := handlers.AllowCredentials()
	originsOk := handlers.AllowedOrigins([]string{os.Getenv("ORIGIN_ALLOWED")})
	methodsOk := handlers.AllowedMethods([]string{"GET", "DELETE", "HEAD", "POST", "PUT", "OPTIONS"})
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homePage).Methods("GET")
	router.HandleFunc("/create/", control.CreateUsers).Methods("POST")
	router.HandleFunc("/read/", control.ReadAllUsers).Methods("GET")
	router.HandleFunc("/update/{uname}/", control.UpdateUsers).Methods("POST")
	router.HandleFunc("/delete/{uname}", control.DeleteUser).Methods("POST")
	log.Fatal(http.ListenAndServe(":4008", handlers.CORS(originsOk, headersOk, methodsOk)(router)))
}
func homePage(w http.ResponseWriter, r *http.Request) {
	myName := "test text"
	fmt.Fprintf(w, "Hello, %s", myName)
}
