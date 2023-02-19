// package main
//
// import (
//
//	"fmt"
//	"net/http"
//	"os"
//
// )
//
//	func findAverage(a []int) float64 {
//		count := 4
//		sum := 0
//		for i := 0; i < count; i++ {
//			sum += (a[i])
//		}
//
//		return float64(sum)
//	}
//
//	func main() {
//		i := []int{5, 6, 7, 8}
//		fmt.Println("AVERAGE", findAverage(i))
//	}
//
// package main
//
// import (
//
//	"net/http"
//	"os"
//
// )
//
//	func indexHandler(w http.ResponseWriter, r *http.Request) {
//		w.Write([]byte("<h1>Hello Everybody!</h1>"))
//	}
//
//	func main() {
//		port := os.Getenv("PORT")
//		if port == "" {
//			port = "3001"
//		}
//
//		mux := http.NewServeMux()
//
//		mux.HandleFunc("/", indexHandler)
//		http.ListenAndServe(":"+port, mux)
//	}
//
// "C://Users//onbol//GolandProjects//awesomeProject//ShopGo//data.db"
package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	// create a database object which can be
	// used to connect with database.
	db, err := sql.Open("mysql", "root:onbolsyn.2004@tcp(0.0.0.0:3306)/sakila")

	// handle error, if any.
	if err != nil {
		panic(err)
	}

	// Here a SQL query is used to return all
	// the data from the table user.
	result, err := db.Query("SELECT actor_id, first_name, last_name FROM sakila.actor")

	// handle error
	if err != nil {
		panic(err)
	}

	// the result object has a method called Next,
	// which is used to iterate through all returned rows.
	for result.Next() {

		var actor_id int
		var first_name string
		var last_name string

		// The result object provided Scan method
		// to read row data, Scan returns error,
		// if any. Here we read id and name returned.
		err = result.Scan(&actor_id, &first_name, &last_name)

		// handle error
		if err != nil {
			panic(err)
		}

		fmt.Println("Id", actor_id, "Name: ", first_name, "Surname: ", last_name)
	}

	// database object has a method Close,
	// which is used to free the resource.
	// Free the resource when the function
	// is returned.
	defer db.Close()
}
