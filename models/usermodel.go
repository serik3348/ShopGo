package models

import (
	"database/sql"
	"server/config"
	"server/entities"
)

type UserModel struct {
	db *sql.DB
}

func NewUserModel() *UserModel {
	conn, err := config.DBConn()
	if err != nil {
		panic(err)
	}
	return &UserModel{
		db: conn,
	}
}

func (u UserModel) Where(user *entities.User, fieldName, fieldValue string) error {
	row, err := u.db.Query("select * from users where "+fieldName+"=? limit 1", fieldValue)
	if err != nil {
		return err
	}
	defer row.Close()

	for row.Next() {
		row.Scan(&user.Id, &user.Nameandsurname, &user.Email, &user.Username, &user.Password)
	}
	return nil
}
func (u UserModel) Create(user entities.User) (int64, error) {
	result, err := u.db.Exec("INSERT INTO users (nameandsurname,email,username,password) values (?,?,?,?)",
		user.Nameandsurname, user.Email, user.Username, user.Password)

	if err != nil {
		return 0, err
	}

	lastInserted, _ := result.LastInsertId()
	return lastInserted, nil
}
