package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"io/ioutil"
	"io"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
	//todos := Todos{
	//	To.do{Name: "Write presentation"},
	//	T.odo{Name: "Host meetup"},
	//}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	fmt.Fprintln(w, "Todo show:", todoId)
}

func TodoCreate(w http.ResponseWriter, r *http.Request) {
	var todo Todo

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &todo); err != nil {
		w.Header().Set("Content-type", "application/json; charset=UTF-8")
		w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
	t := RepoCreateTodo(todo)
	w.Header().Set("Content-type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}

//func GetUser(w http.ResponseWriter, r *http.Request) {
//	json.NewEncoder(w).Encode(users)
//}


//func UpdateUser(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
//	w.WriteHeader(http.StatusOK)
//
//	if err := json.NewEncoder(w).Encode(users); err != nil {
//		panic(err)
//	}
//
//}
//
//func CreateUser(w http.ResponseWriter, r *http.Request) {
//	params := mux.Vars(r)
//	var user User
//	_ = json.NewDecoder(r.Body).Decode(&user)
//	Id := params["id"]
//	fmt.Fprintln(w, "show user: ", Id)
//	users = append(users, user)
//	json.NewEncoder(w).Encode(users)
//}

func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var user User
	Id := params["Id"]
	fmt.Fprintln(w, "user get:", Id)
	json.NewEncoder(w).Encode(&user)
}

func GetPeople(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(users)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)
	Id := params["id"]
	fmt.Fprintln(w, "user post:", Id)
	users = append(users, user)
	json.NewEncoder(w).Encode(users)
}

//func DeleteUser(w http.ResponseWriter, r *http.Request) {
//	params := mux.Vars(r)
//	for index, item := range users {
//		if item.Id == params["id"] {
//			users = append(users[:index], users[index+1:]...)
//			break
//		}
//	}
//	json.NewEncoder(w).Encode(users)
//}

