package controllers

import (
	"html/template"
	"net/http"
	"server/config"
	"server/models"
)

func IndexProd(w http.ResponseWriter, r *http.Request) {
	session, _ := config.Store.Get(r, config.SESSION_ID)
	var productModel models.ProductModel
	products, _ := productModel.FindAll()
	data := map[string]interface{}{
		"products":       products,
		"nameandsurname": session.Values["nameandsurname"],
	}
	tmp, _ := template.ParseFiles("midterm1/views/product/index.html")
	tmp.Execute(w, data)
}

func Filter(w http.ResponseWriter, r *http.Request) {
	session, _ := config.Store.Get(r, config.SESSION_ID)
	//query := r.URL.Query()
	//filter, _ := query.Get("filter")
	////query := r.URL.Query()
	////id, _ := strconv.ParseInt(query.Get("id"), 10, 64)
	fromorder := r.FormValue("fromorder")
	toorder := r.FormValue("toorder")

	var productModel models.ProductModel
	products, _ := productModel.Filter(fromorder, toorder)

	data := map[string]interface{}{
		"products":       products,
		"nameandsurname": session.Values["nameandsurname"],
	}

	tmp, _ := template.ParseFiles("midterm1/views/product/filter.html")
	tmp.Execute(w, data)

}

func Search(w http.ResponseWriter, r *http.Request) {

	//query := r.URL.Query()
	//id, _ := strconv.ParseInt(query.Get("id"), 10, 64)
	var name = r.FormValue("name")

	var productModel models.ProductModel
	product, _ := productModel.Searcher(name)
	session, _ := config.Store.Get(r, config.SESSION_ID)
	data := map[string]interface{}{
		"name":           name,
		"product":        product,
		"nameandsurname": session.Values["nameandsurname"],
	}
	tmp, _ := template.ParseFiles("midterm1/views/product/search.html")
	tmp.Execute(w, data)

}
