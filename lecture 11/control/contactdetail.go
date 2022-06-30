package control

import (
	"encoding/json"
	"fmt"
	"model/model"

	"net/http"

	"github.com/gorilla/mux"
)

func CreateContactDetail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	var newContact model.ContactDetail
	_ = json.NewDecoder(r.Body).Decode(&newContact)
	// fmt.Println(newContact.Fname)
	// fmt.Println(newContact.Lname)
	// fmt.Println(newContact.Age)
	model.DetailContact = append(model.DetailContact, newContact)
	w.WriteHeader(201)
	fmt.Fprintf(w, "Added")

}

func ReadContactDetail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	//var newContact model.ContactDetail
	// for i := 0; i < len(model.DetailContact); i++ {
	// 	if model.DetailContact[i].IsActive == true {
	// 		activeuser = append(activeuser, model.DetailContact[i])
	// 	}
	// }

	json.NewEncoder(w).Encode(model.DetailContact)
	w.WriteHeader(200)
}

// func UpdateUser(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
// 	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
//
// 	json.NewEncoder(w).Encode(view.AllUser[index])
// 	w.WriteHeader(200)

// }

func UpdateContactDetail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	var params = mux.Vars(r)
	var types string = params["type"]
	// fmt.Println(uname)
	var index = -1
	for i := 0; i < len(model.AllContact); i++ {
		if model.DetailContact[i].Type == types {
			index = i
			break
		}
	}
	if index == -1 {
		w.WriteHeader(404)
		fmt.Fprintf(w, "%s Not Found", types)
		return
	} else {
		var updateuser model.ContactDetail

		_ = json.NewDecoder(r.Body).Decode(&updateuser)
		// fmt.Println(updateuser.Fname)
		// fmt.Println(updateuser.Lname)
		if updateuser.Type != "" {
			model.DetailContact[index].Type = updateuser.Type
		}
		if updateuser.Number != 0 {
			model.DetailContact[index].Number = updateuser.Number
		}
	}
	w.WriteHeader(201)
	fmt.Fprintf(w, "Added")
}

//
// _ = json.NewDecoder(r.Body).Decode(&updateuser)
// view.AllUser = append(view.AllUser, updateuser)
// w.WriteHeader(201)
// fmt.Fprintf(w, "Added")

func DeleteContactDetail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	var params = mux.Vars(r)
	var types string = params["type"]
	var index = -1
	for i := 0; i < len(model.DetailContact); i++ {
		if model.DetailContact[i].Type == types {
			index = i
			break
		}
	}
	if index == -1 {
		w.WriteHeader(404)
		fmt.Fprintf(w, "%s Not Found", types)
		return
	} else {
		model.DetailContact[index].IsActive = false
	}
	json.NewEncoder(w).Encode(model.DetailContact[index])
	w.WriteHeader(200)

}

// func GetUserDetails(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
// 	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
// 	var params = mux.Vars(r)
// 	var fullname string = params["fullname"]
// 	var index = -1
// 	for i := 0; i < len(view.AllUser); i++ {
// 		if view.AllUser[i].FullName == fullname {
// 			index = i
// 			break
// 		}
// 	}
// 	if index == -1 {
// 		w.WriteHeader(404)
// 		fmt.Fprintf(w, "%s Not Found", fullname)
// 		return
// 	}
// 	json.NewEncoder(w).Encode(view.AllUser[index])
// 	w.WriteHeader(200)
// }
