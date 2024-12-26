package models

import (
	"errors"

	"airbnb.com/airbnb/db"
)

type Feature struct {
	Id        int64
	Place_id  int64
	Space     int64 `binding:"required"`
	Bathrooms int64 `binding:"required"`
	Bedrooms  int64 `binding:"required"`
	Kitchen   int64 `binding:"required"`
	Capacity  int64 `binding:"required"`
}

func GetPlaceFeatures(id int64) (*Feature, error) {
	query := `SELECT * FROM features WHERE place_id = ?`

	row := db.DB.QueryRow(query, id)

	features := Feature{}

	err := row.Scan(&features.Id, &features.Place_id, &features.Space, &features.Bathrooms, &features.Bedrooms, &features.Kitchen, &features.Capacity)

	if err != nil {
		return nil, errors.New("cant get features")
	}

	return &features, nil

}

func (feature *Feature) AddNewPlaceFeature() (int64, error) {
	query := `INSERT INTO features (place_id , space , bathrooms , bedrooms , kitchen , capacity) VALUES (?,?,?,?,?,?)`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return -1, errors.New("cant add new feature")
	}

	resutl, err := stmt.Exec(&feature.Place_id, &feature.Space, &feature.Bathrooms, &feature.Bedrooms, &feature.Kitchen, &feature.Capacity)

	if err != nil {
		return -1, errors.New("cant add new feature")
	}

	id, err := resutl.LastInsertId()

	return id, err

}
