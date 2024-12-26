package models

import (
	"errors"

	"airbnb.com/airbnb/db"
)

type User struct {
	Id        int64
	FirstName string `binding:"required"`
	LastName  string `binding:"required"`
	Email     string `binding:"required"`
	Password  string `binding:"required"`
}

type UserLogin struct {
	Id        int64
	FirstName string
	LastName  string
	Email     string `binding:"required"`
	Password  string `binding:"required"`
}

func (u *User) Signup() (int64, error) {
	query := `INSERT INTO users (first_name , last_name , email , password) VALUES (?,?,?,?)`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return -1, errors.New("cant signup user")
	}

	result, err := stmt.Exec(u.FirstName, u.LastName, u.Email, u.Password)

	if err != nil {
		return -1, errors.New("cant signup user")
	}

	id, err := result.LastInsertId()

	return id, err
}

func (u *UserLogin) ValidateUserCreadentials() (string, error) {
	query := `SELECT password , id FROM users WHERE email = ?`
	var retrivedPassword string
	row := db.DB.QueryRow(query, &u.Email)

	err := row.Scan(&retrivedPassword, &u.Id)

	return retrivedPassword, err

}
