package models

import (
	"database/sql"
	"server/config"
	"server/entities"
)

type CommentModel struct {
	db *sql.DB
}

func NewCommentModel() *CommentModel {
	conn, err := config.DBConn()
	if err != nil {
		panic(err)
	}
	return &CommentModel{
		db: conn,
	}
}

func (*CommentModel) Find(name string) ([]entities.Comment, error) {
	db, err := config.DBConn()
	if err != nil {
		return []entities.Comment{}, err
	} else {
		rows, err2 := db.Query("select * from comment where objectname=? ", name)
		if err2 != nil {
			return []entities.Comment{}, err2
		} else {
			var comments []entities.Comment
			var comment entities.Comment
			for rows.Next() {
				rows.Scan(&comment.Id, &comment.Username, &comment.ObjectName, &comment.Text, &comment.Date)
				comments = append(comments, comment)
			}
			return comments, nil
		}
	}
}

func (u CommentModel) Create(comment entities.Comment) (int64, error) {
	result, err := u.db.Exec("INSERT INTO comment (username,objectname,text,date) values (?,?,?,CURRENT_DATE)",
		comment.Username, comment.ObjectName, comment.Text)
	if err != nil {
		return 0, err
	}

	lastInserted, _ := result.LastInsertId()
	return lastInserted, nil
}
