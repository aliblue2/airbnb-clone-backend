package models

import (
	"errors"
	"time"

	"airbnb.com/airbnb/db"
)

type Comment struct {
	Id         int64
	Place_id   int64
	User_id    int64
	Content    string `binding:"required"`
	Created_at time.Time
	Rate       int64 `binding:"required"`
}

func GetCommentsPlace(id int64) (*[]Comment, error) {
	query := `SELECT * FROM comments WHERE place_id = ?`

	rows, err := db.DB.Query(query, id)

	if err != nil {
		return nil, errors.New("cant get all comments")
	}
	comments := []Comment{}
	for rows.Next() {
		tempComment := Comment{}

		err := rows.Scan(&tempComment.Id, &tempComment.Place_id, &tempComment.User_id, &tempComment.Content, &tempComment.Created_at, &tempComment.Rate)

		if err != nil {
			return nil, errors.New("cant get all comments")
		}

		comments = append(comments, tempComment)
	}

	return &comments, nil

}

func (comment *Comment) AddNewComments() (int64, error) {

	query := `INSERT INTO comments (place_id , user_id , content , created_at , rate) VALUES (?,?,?,?,?)`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return -1, errors.New("cant add new comment")
	}

	result, err := stmt.Exec(comment.Place_id, comment.User_id, comment.Content, comment.Created_at, comment.Rate)

	if err != nil {
		return -1, errors.New("cant add new comment")
	}

	id, err := result.LastInsertId()

	return id, err

}
