package models

import (
	"errors"

	"airbnb.com/airbnb/db"
)

type Place struct {
	Id          int64
	OwnerId     int64
	Name        string `binding:"required" json:"name"`
	Description string `binding:"required" json:"description"`
	Price       string `binding:"required" json:"price"`
	State       string `binding:"required" json:"state"`
	City        string `binding:"required" json:"city"`
	Banner      string `binding:"required" json:"banner"`
	Images      string
	Rate        string
	Latitude    string
	Longitude   string
}

func GetAllPlaces() ([]Place, error) {
	places := []Place{}
	query := `SELECT * FROM places`

	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, errors.New("cant get all places")
	}

	for rows.Next() {
		tempPlace := Place{}
		err := rows.Scan(&tempPlace.Id, &tempPlace.OwnerId, &tempPlace.Name, &tempPlace.Description, &tempPlace.Price, &tempPlace.State, &tempPlace.City, &tempPlace.Banner, &tempPlace.Images, &tempPlace.Rate, &tempPlace.Latitude, &tempPlace.Longitude)
		if err != nil {
			return nil, errors.New("cant get all places")
		}
		places = append(places, tempPlace)
	}

	return places, nil

}

func GetPlaceById(id int64) (*Place, error) {
	query := `SELECT * FROM places WHERE id = ?`

	row := db.DB.QueryRow(query, id)

	place := Place{}

	err := row.Scan(&place.Id, &place.OwnerId, &place.Name, &place.Description, &place.Price, &place.State, &place.City, &place.Banner, &place.Images, &place.Rate, &place.Latitude, &place.Longitude)

	if err != nil {
		return nil, errors.New("cant get place")
	}

	return &place, nil
}

func (place *Place) CreateNewPlace() (int64, error) {
	query := `INSERT INTO places (owner_id , name , description , price, state, city , banner , images , rate , latitude , longitude) VALUES (?,?,?,?,?,?,?,?,?,?,?)`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return -1, errors.New("cant create new place")
	}

	result, err := stmt.Exec(place.OwnerId, place.Name, place.Description, place.Price, place.State, place.City, place.Banner, place.Images, place.Rate, place.Latitude, place.Longitude)

	if err != nil {
		return -1, errors.New("cant create new place")
	}

	id, err := result.LastInsertId()

	return id, err

}

func DeletePlace(id int64) error {
	query := `DELETE FROM places WHERE id = ? `

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return errors.New("cant delete with this item")
	}

	result, err := stmt.Exec(id)
	if err != nil {
		return errors.New("cant delete with this item")
	}

	rowsEffected, err := result.RowsAffected()

	if rowsEffected <= 0 {
		return errors.New("cant delete item")
	}

	return nil

}

func (place *Place) EditPlaceById(id int64) error {
	query := `UPDATE places SET
	owner_id = ?, 
	name = ? ,
	description = ?, 
	price = ?,
	state = ?,
	city = ? ,
	banner = ? ,
	images = ? ,
	rate = ? ,
	latitude = ? ,
	longitude = ? 
	WHERE id = ?
	`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return errors.New("cant update place item")
	}

	resutl, err := stmt.Exec(&place.OwnerId, &place.Name, &place.Description, &place.Price, &place.State, &place.City, &place.Banner, &place.Images, &place.Rate, &place.Latitude, &place.Longitude, id)

	if err != nil {
		return errors.New("cant update place item")
	}

	rowsEffect, err := resutl.RowsAffected()

	if rowsEffect <= 0 {
		return errors.New("no places updated.!")
	}

	return nil

}
