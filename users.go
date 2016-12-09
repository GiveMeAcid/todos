package main

type User struct {
	Login string `json:"login"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type Users []User

func (u *User) SetPassword(plainPassword string) {
	u.Password = plainPassword
}

//type ChangeSettings struct {
//	Login string `json:"login"`
//	Name string `json:"name"`
//	Email string `json:"email"`
//	Password string `json:"password"`
//}