package control

import (
	"encoding/json"
	"fmt"
	"model/model"

	"net/http"

	"github.com/gorilla/mux"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	var newUser model.User
	_ = json.NewDecoder(r.Body).Decode(&newUser)
	fmt.Println(newUser.Fname)
	fmt.Println(newUser.Lname)

	model.CreateUser(newUser.UniqueID, newUser.Fname, newUser.Lname, newUser.Password, newUser.IsAdmin)

	w.WriteHeader(201)
	fmt.Fprintf(w, "User Added")

}

func ReadUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	IsSessionValid := model.IsValidCoockie(w, r)

	if !IsSessionValid {
		fmt.Println("Session expired")
	} else {
		var ActiveUser []model.User
		for i := 0; i < len(model.AllUser); i++ {

			if model.AllUser[i].IsActive {
				ActiveUser = append(ActiveUser, model.AllUser[i])
			}
		}

		json.NewEncoder(w).Encode(ActiveUser)
		w.WriteHeader(200)
	}

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	var params = mux.Vars(r)
	var UniqueId string = params["unique"]
	fmt.Println(UniqueId)
	var index = -1
	for i := 0; i < len(model.AllUser); i++ {
		fmt.Println(model.AllUser[i].UniqueID, "      ::", UniqueId)
		if model.AllUser[i].UniqueID == UniqueId {

			index = i
			break
		}
	}
	if index == -1 {
		w.WriteHeader(404)
		fmt.Fprintf(w, "%s Not Found", UniqueId)
		return
	} else {
		var updateuser model.User

		_ = json.NewDecoder(r.Body).Decode(&updateuser)

		if updateuser.Fname != "" {
			model.AllUser[index].Fname = updateuser.Fname
		}
		if updateuser.Lname != "" {
			model.AllUser[index].Lname = updateuser.Lname
		}
	}
	w.WriteHeader(201)
	fmt.Fprintf(w, "User Updated")
}

func DeleteUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "delete fUNCTION")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	var params = mux.Vars(r)
	var UniqueId string = params["unique"]
	fmt.Println(UniqueId)
	var index = -1
	for i := 0; i < len(model.AllUser); i++ {
		fmt.Println(model.AllUser[i].UniqueID, "      ::", UniqueId)
		if model.AllUser[i].UniqueID == UniqueId {

			index = i
			break
		}
	}
	if index == -1 {
		w.WriteHeader(404)
		fmt.Fprintf(w, "%s Not Found", UniqueId)
		return
	} else {
		model.AllUser[index].IsActive = false
	}
	w.WriteHeader(201)
	fmt.Fprintf(w, "User Deleted")

}
