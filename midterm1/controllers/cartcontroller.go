package controllers

import (
	"encoding/json"
	"html/template"
	"net/http"
	"server/config"
	"server/entities"
	"server/models"
	"strconv"
)

func IndexCart(w http.ResponseWriter, r *http.Request) {
	session, _ := config.Store.Get(r, config.SESSION_ID)
	strCart := session.Values["cart"].(string)
	var cart []entities.Item
	json.Unmarshal([]byte(strCart), &cart)

	data := map[string]interface{}{
		"cart":           cart,
		"total":          total(cart),
		"nameandsurname": session.Values["nameandsurname"],
	}

	tmp, _ := template.New("index.html").Funcs(template.FuncMap{
		"total": func(item entities.Item) float64 {
			return item.Product.Price * float64(item.Quantity)
		}}).ParseFiles("midterm1/views/cart/index.html")
	tmp.Execute(w, data)
}

func Buy(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	id, _ := strconv.ParseInt(query.Get("id"), 10, 64)
	var productModel models.ProductModel
	product, _ := productModel.Find(id)
	session, _ := config.Store.Get(r, config.SESSION_ID)
	cart := session.Values["cart"]

	if cart == nil {
		var cart []entities.Item
		cart = append(cart, entities.Item{
			Product:  product,
			Quantity: 1,
		})
		bytesCart, _ := json.Marshal(cart)
		session.Values["cart"] = string(bytesCart)
	} else {
		strCart := session.Values["cart"].(string)
		var cart []entities.Item
		json.Unmarshal([]byte(strCart), &cart)

		index := exists(id, cart)
		if index == -1 {
			cart = append(cart, entities.Item{
				Product:  product,
				Quantity: 1,
			})
		} else {
			cart[index].Quantity++
		}
		bytesCart, _ := json.Marshal(cart)
		session.Values["cart"] = string(bytesCart)
	}

	session.Save(r, w)
	http.Redirect(w, r, "/cart", http.StatusSeeOther)
}

func Remove(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	id, _ := strconv.ParseInt(query.Get("id"), 10, 64)
	session, _ := config.Store.Get(r, config.SESSION_ID)
	strCart := session.Values["cart"].(string)

	var cart []entities.Item
	json.Unmarshal([]byte(strCart), &cart)

	index := exists(id, cart)
	cart = remove(cart, index)

	bytesCart, _ := json.Marshal(cart)
	session.Values["cart"] = string(bytesCart)
	session.Save(r, w)
	http.Redirect(w, r, "/cart", http.StatusSeeOther)
}

func exists(id int64, cart []entities.Item) int {
	for i := 0; i < len(cart); i++ {
		if cart[i].Product.Id == id {
			return i
		}
	}
	return -1
}

func total(cart []entities.Item) float64 {
	var s float64 = 0
	for _, item := range cart {
		s += item.Product.Price * float64(item.Quantity)
	}
	return s
}

func remove(cart []entities.Item, index int) []entities.Item {
	copy(cart[index:], cart[index+1:])
	return cart[:len(cart)-1]
}
