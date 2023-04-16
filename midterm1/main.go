package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"server/controllers"

	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("midterm1/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/login", controllers.Login)
	http.HandleFunc("/logout", controllers.Logout)
	http.HandleFunc("/register", controllers.Register)
	http.HandleFunc("/product", controllers.IndexProd)
	http.HandleFunc("/product/index", controllers.IndexProd)
	http.HandleFunc("/cart", controllers.IndexCart)
	http.HandleFunc("/cart/index", controllers.IndexCart)
	http.HandleFunc("/cart/buy", controllers.Buy)
	http.HandleFunc("/cart/remove", controllers.Remove)
	http.HandleFunc("/product/search", controllers.Search)
	http.HandleFunc("/product/filter", controllers.Filter)

	fmt.Println("Server address : http://localhost:3000")
	http.ListenAndServe(":3000", nil)

}
