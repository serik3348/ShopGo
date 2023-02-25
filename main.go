package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"

	"net/http"
)

func main() {

	http.HandleFunc("/", Index)
	http.HandleFunc("/login", Login)
	http.HandleFunc("/logout", Logout)
	http.HandleFunc("/register", Register)
	fmt.Println("Server address : http://localhost:3000")
	http.ListenAndServe(":3000", nil)
}
