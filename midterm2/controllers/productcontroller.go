package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"server/config"
	"server/entities"
	"server/models"
	"strconv"
)

func IndexProd(w http.ResponseWriter, r *http.Request) {
	session, _ := config.Store.Get(r, config.SESSION_ID)
	var productModel models.ProductModel
	//var rankModel models.RankModel

	products, _ := productModel.FindAll()
	//var rank []float64
	//for i:=0;i< len(products);i++{
	//	rank[i]=rankModel.AvgRank(products[i].Name)
	//}
	data := map[string]interface{}{
		"products": products,

		"nameandsurname": session.Values["nameandsurname"],
	}
	tmp, _ := template.ParseFiles("midterm2/views/product/index.html")
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

	tmp, _ := template.ParseFiles("midterm2/views/product/filter.html")
	tmp.Execute(w, data)

}

func Search(w http.ResponseWriter, r *http.Request) {

	//query := r.URL.Query()
	//id, _ := strconv.ParseInt(query.Get("id"), 10, 64)
	var name = r.FormValue("name")

	var productModel models.ProductModel
	var commentModel models.CommentModel
	var rankModel models.RankModel

	product, _ := productModel.Searcher(name)
	comments, _ := commentModel.Find(name)

	var sum float64

	rank := rankModel.AvgRank(name)
	for i := 0; i < len(rank); i++ {
		sum += rank[i].Ranking
	}
	answer := (sum) / (float64(len(rank)))

	session, _ := config.Store.Get(r, config.SESSION_ID)

	data := map[string]interface{}{

		"product":        product,
		"answer":         answer,
		"comments":       comments,
		"nameandsurname": session.Values["nameandsurname"],
		"username":       session.Values["username"],
	}
	tmp, _ := template.ParseFiles("midterm2/views/product/search.html")
	tmp.Execute(w, data)

}

func Rank(w http.ResponseWriter, r *http.Request) {

	//session, _ := config.Store.Get(r, config.SESSION_ID)

	r.ParseForm()
	rankname := r.Form.Get("rankname")
	ranking, _ := strconv.ParseFloat(r.Form.Get("ranking"), 64)

	fmt.Println(rankname, ranking)
	var rankmodel models.RankModel

	rank := entities.Rank{
		ObjectName: rankname,
		Ranking:    ranking,
	}
	rankmodel.GiveRank(rank)
	fmt.Println(rank)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
