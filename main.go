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
	_ "github.com/go-sql-driver/mysql"
	authcontroller "github.com/serik3348/ShopGo/controllers"
	"net/http"
)

func main() {
	//db, err := sql.Open("mysql", "root:onbolsyn.2004@tcp(0.0.0.0:3306)/golang")
	//
	//if err != nil {
	//	panic(err)
	//}
	//_, err = db.Exec("CREATE TABLE user(id INT NOT NULL, name VARCHAR(20), email VARCHAR(30), PRIMARY KEY (ID));")
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Print("Successfully Created\n")
	//defer db.Close()
	http.HandleFunc("/", authcontroller.Index)
}
