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
