package controllers

import (
	"net/http"
	"server/config"
	"server/entities"
	"server/models"
)

var commentModel = models.NewCommentModel()

func WriteComment(w http.ResponseWriter, r *http.Request) {

	session, _ := config.Store.Get(r, config.SESSION_ID)
	//query := r.URL.Query()
	//id, _ := strconv.ParseInt(query.Get("id"), 10, 64)

	var productModel models.ProductModel

	//var objectname = r.FormValue("object_name")
	r.ParseForm()
	id := r.Form.Get("id")
	username := session.Values["username"].(string)
	//num, _ := strconv.ParseInt(id, 10, 64)
	product, _ := productModel.Searcher(id)

	comment := entities.Comment{
		Username:   username,
		ObjectName: product.Name,
		Text:       r.Form.Get("text"),
	}
	commentModel.Create(comment)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
