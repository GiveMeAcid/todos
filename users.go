package main

type User struct {
	Id int `json:"id"`
	Login string `json:"login"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type Users []User