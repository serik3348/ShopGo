package config

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func DBConn() (db *sql.DB, err error) {
	db, err = sql.Open("mysql", "root:onbolsyn.2004@tcp(0.0.0.0:3306)/golang")
	return db, err
}
