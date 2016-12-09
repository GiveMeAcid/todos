package main

import (
	"log"
	"net/http"
	"github.com/jinzhu/gorm"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

func main() {

	db, err := gorm.Open("postgres", "host=localhost user=postgres dbname=testmuxdb sslmode=disable password=31780")

	if err != nil {
		fmt.Printf("Database opening error -->%v\n", err)
		panic("Database error")
	}
	defer db.Close()

	db.AutoMigrate(&User{})

	db.SingularTable(true)

	router := NewRouter()
	//router.HandleFunc("/users/{login}", UpdateUser).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8080", router))
}