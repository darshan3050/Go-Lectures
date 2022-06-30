package control

import (
	"encoding/json"
	"fmt"

	"model/view"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	var newUser view.User
	_ = json.NewDecoder(r.Body).Decode(&newUser)
	fmt.Println(newUser.Fname)
	fmt.Println(newUser.Lname)
	fmt.Println(newUser.Age)
	view.AllUser = append(view.AllUser, newUser)
	w.WriteHeader(201)
	fmt.Fprintf(w, "Added")

}

func ReadAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	var activeuser []view.User
	for i := 0; i < len(view.AllUser); i++ {
		if view.AllUser[i].IsActive == true {
			activeuser = append(activeuser, view.AllUser[i])
		}
	}

	json.NewEncoder(w).Encode(activeuser)
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

func UpdateUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	var params = mux.Vars(r)
	var uname string = params["uname"]
	fmt.Println(uname)
	var index = -1
	for i := 0; i < len(view.AllUser); i++ {
		if view.AllUser[i].Unique == uname {
			index = i
			break
		}
	}
	if index == -1 {
		w.WriteHeader(404)
		fmt.Fprintf(w, "%s Not Found", uname)
		return
	} else {
		var updateuser view.User

		_ = json.NewDecoder(r.Body).Decode(&updateuser)
		fmt.Println(updateuser.Fname)
		fmt.Println(updateuser.Lname)
		if updateuser.Fname != "" {
			view.AllUser[index].Fname = updateuser.Fname
		}
		if updateuser.Lname != "" {
			view.AllUser[index].Lname = updateuser.Lname
		}
		if updateuser.Age != 0 {
			view.AllUser[index].Age = updateuser.Age
		}
		if updateuser.Location != "" {
			view.AllUser[index].Location = updateuser.Location
		}
		if updateuser.Number != 0 {
			view.AllUser[index].Number = updateuser.Number
		}
		w.WriteHeader(201)
		fmt.Fprintf(w, "Added")
	}
	//
	// _ = json.NewDecoder(r.Body).Decode(&updateuser)
	// view.AllUser = append(view.AllUser, updateuser)
	// w.WriteHeader(201)
	// fmt.Fprintf(w, "Added")

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	var params = mux.Vars(r)
	var name string = params["name"]
	var index = -1
	for i := 0; i < len(view.AllUser); i++ {
		if view.AllUser[i].Fname == name {
			index = i
			break
		}
	}
	if index == -1 {
		w.WriteHeader(404)
		fmt.Fprintf(w, "%s Not Found", name)
		return
	} else {
		view.AllUser[index].IsActive = false
	}
	json.NewEncoder(w).Encode(view.AllUser[index])
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
