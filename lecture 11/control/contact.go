package control

import (
	"encoding/json"
	"fmt"
	"model/model"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateContacts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	var newUser model.Contact
	_ = json.NewDecoder(r.Body).Decode(&newUser)
	fmt.Println(newUser.Fname)
	fmt.Println(newUser.Lname)

	model.AllContact = append(model.AllContact, newUser)
	w.WriteHeader(201)
	fmt.Fprintf(w, "Added")

}

func ReadAllContacts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	var activeuser []model.Contact
	for i := 0; i < len(model.AllContact); i++ {
		if model.AllContact[i].IsActive {
			activeuser = append(activeuser, model.AllContact[i])
		}
	}

	json.NewEncoder(w).Encode(activeuser)
	w.WriteHeader(200)
}

func UpdateContacts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	var params = mux.Vars(r)
	var uname string = params["uname"]
	fmt.Println(uname)
	var index = -1
	for i := 0; i < len(model.AllContact); i++ {
		if model.AllContact[i].Fname == uname {
			index = i
			break
		}
	}
	if index == -1 {
		w.WriteHeader(404)
		fmt.Fprintf(w, "%s Not Found", uname)
		return
	} else {
		var updateuser model.Contact

		_ = json.NewDecoder(r.Body).Decode(&updateuser)
		fmt.Println(updateuser.Fname)
		fmt.Println(updateuser.Lname)
		if updateuser.Fname != "" {
			model.AllContact[index].Fname = updateuser.Fname
		}
		if updateuser.Lname != "" {
			model.AllContact[index].Lname = updateuser.Lname
		}

		w.WriteHeader(201)
		fmt.Fprintf(w, "Added")
	}

}

func DeleteContacts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	var params = mux.Vars(r)
	var name string = params["name"]
	var index = -1
	for i := 0; i < len(model.AllContact); i++ {
		if model.AllContact[i].Fname == name {
			index = i
			break
		}
	}
	if index == -1 {
		w.WriteHeader(404)
		fmt.Fprintf(w, "%s Not Found", name)
		return
	} else {
		model.AllContact[index].IsActive = false
	}
	json.NewEncoder(w).Encode(model.AllContact[index])
	w.WriteHeader(200)

}
