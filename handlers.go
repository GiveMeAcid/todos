package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"io/ioutil"
	"io"
	"github.com/jinzhu/gorm"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

func TodoShow(w http.ResponseWriter, r *http.Request) {

	//vars := mux.Vars(r)
	//todoId := vars["Id"]
	//fmt.Fprintln(w, "Todo show:", todoId)

	t := RepoFindTodo(currentId)
	w.Header().Set("Content-type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
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

func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range users {
		if item.Login == params["login"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&User{})
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(users)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var (
		DB *gorm.DB
		user User
		//email = r.PostFormValue("email")
		password = r.PostFormValue("password")
		//name = r.PostFormValue("name")
	)
	//params := mux.Vars(r)
	_ = json.NewDecoder(r.Body).Decode(&user)
	DB.Where(user).First(&user)
	user.SetPassword(password)
	//user.Login = params["login"]
	users = append(users, user)
	json.NewEncoder(w).Encode(user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range users {
		if item.Login == params["login"] {
			users = append(users[:index], users[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(users)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	vars := mux.Vars(r)
	login := vars["login"]
	for _, p := range users {
		if p.Login == login {
			user = p
		}
	}

	w.Header().Set("Content-Type", "application/json")
	if user.Login != "" {
		payload, _ := json.Marshal(user)
		w.Write([]byte(payload))
	} else {
		w.Write([]byte("User Not Found"))
	}
	json.NewEncoder(w).Encode(user)
}

