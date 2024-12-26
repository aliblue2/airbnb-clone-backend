package models

import (
	"errors"

	"airbnb.com/airbnb/db"
)

type Owner struct {
	Id        int64
	FirstName string `binding:"required"`
	LastName  string `binding:"required"`
	Phone     string `binding:"required"`
	Password  string `binding:"required"`
}

type OwnerLogin struct {
	Id        int64
	FirstName string
	LastName  string
	Phone     string `binding:"required"`
	Password  string `binding:"required"`
}

func (o *Owner) Signup() (int64, error) {
	query := `INSERT INTO owners (first_name , last_name , phone , password) VALUES (?,?,?,?)`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return -1, errors.New("cant signup user")
	}

	result, err := stmt.Exec(&o.FirstName, &o.LastName, &o.Phone, &o.Password)

	if err != nil {
		return -1, errors.New("cant signup user")
	}

	id, err := result.LastInsertId()

	return id, err
}

func (o *OwnerLogin) ValidateUserCreadentials() (string, error) {
	query := `SELECT password , id FROM owners WHERE phone = ?`
	var retrivedPassword string
	row := db.DB.QueryRow(query, &o.Phone)

	err := row.Scan(&retrivedPassword, &o.Id)

	return retrivedPassword, err

}
