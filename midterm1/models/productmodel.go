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
func (*ProductModel) Searcher(name string) (entities.Product, error) {
	db, err := config.DBConn()
	if err != nil {
		return entities.Product{}, err
	} else {
		rows, err2 := db.Query("select * from product where name = ?", name)
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

func (*ProductModel) Decrease(id int64) {
	db, err := config.DBConn()
	if err == nil {

		row, err := db.Query("UPDATE product SET quantity=quantity-1 WHERE id=?", id)
		if err != nil {
			panic(err.Error())
		}
		defer row.Close()

		// Execute the SQL statement

	}
}
func (*ProductModel) Increase(id int64, num int64) {
	db, err := config.DBConn()
	if err == nil {

		row, err := db.Query("UPDATE product SET quantity=quantity+? WHERE id=?", num, id)
		if err != nil {
			panic(err.Error())
		}
		defer row.Close()

		// Execute the SQL statement

	}
}
