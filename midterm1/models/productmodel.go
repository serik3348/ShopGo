package models

import (
	"server/config"
	"server/entities"
)

type ProductModel struct {
}

func (*ProductModel) FindAll() ([]entities.Product, error) {
	db, err := config.DBConn()
	if err != nil {
		return nil, err
	} else {
		rows, err2 := db.Query("select * from product")
		if err2 != nil {
			return nil, err2
		} else {
			var products []entities.Product
			var product entities.Product
			for rows.Next() {
				rows.Scan(&product.Id, &product.Name, &product.Price, &product.Quantity, &product.Photo)
				products = append(products, product)
			}
			return products, nil
		}
	}
}

func (*ProductModel) Find(id int64) (entities.Product, error) {
	db, err := config.DBConn()
	if err != nil {
		return entities.Product{}, err
	} else {
		rows, err2 := db.Query("select * from product where id = ?", id)
		if err2 != nil {
			return entities.Product{}, err2
		} else {

			var product entities.Product
			for rows.Next() {
				rows.Scan(&product.Id, &product.Name, &product.Price, &product.Quantity, &product.Photo)

			}
			return product, nil
		}
	}
}
