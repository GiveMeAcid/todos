package main

import "fmt"

var currentId int
var currentLogin string
var todos Todos
var users Users

func init() {
	RepoCreateTodo(Todo{Name: "Write presentation"})
	RepoCreateTodo(Todo{Name: "Host meetup"})
	CreateNewUser(User{Name: "userlol"})
	CreateNewUser(User{Name: "user"})
	CreateNewUser(User{Name: "test", Login: "testlogin", Password: "testpassword", Email: "test@mail.com"})
}

func RepoFindTodo(id int) Todo{
	for _, t := range todos {
		if t.Id == id {
			return t
		}
	}
	//return empty todos if not created
	return Todo{}
}

func RepoCreateTodo(t Todo) Todo{
	currentId += 1
	t.Id = currentId
	todos = append(todos, t)
	return t
}

func CreateNewUser(t User) User{
	t.Login = currentLogin
	users = append(users, t)
	return t
}

func UserFind(id int) User{
	for _, t := range users {
		if t.Login == currentLogin {
			return t
		}
	}
	return User{}
}

func RepoDestroyTodo(id int) error{
	for i, t := range todos {
		if t.Id == id {
			todos = append(todos[:i], todos[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Todo with id of %d to delete", id)
}