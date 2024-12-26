package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDb() {
	db, err := sql.Open("sqlite3", "api.db")

	if err != nil {
		panic(err)
	}

	DB = db

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	createTables()
}

func createTables() {

	createPlacesTable := `CREATE TABLE IF NOT EXISTS places (
		id INTEGER PRIMARY KEY AUTOINCREMENT, 
		owner_id INTEGER NOT NULL,
		name TEXT NOT NULL, 
		description TEXT NOT NULL, 
		price TEXT NOT NULL, 
		state TEXT NOT NULL, 
		city TEXT NOT NULL, 
		banner TEXT NOT NULL, 
		images TEXT NOT NULL,
		rate INTEGR NOT NULL, 
		latitude TEXT NOT NULL,
		longitude TEXT NOT NULL,
		FOREIGN KEY (owner_id) REFERENCES owners (id)
	)`
	_, err := DB.Exec(createPlacesTable)

	if err != nil {
		panic(err)
	}

	createCommentsTable := `CREATE TABLE IF NOT EXISTS comments (
		id INTEGER PRIMARY KEY AUTOINCREMENT, 
		place_id INTEGER NOT NULL,
		user_id INTEGER NOT NULL, 
		content TEXT NOT NULL, 
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
		rate INTEGER NOT NULL, 
		FOREIGN KEY (place_id) REFERENCES  places (id)
 	)`
	_, err = DB.Exec(createCommentsTable)

	if err != nil {
		panic(err)
	}

	createFeaturesTable := `CREATE TABLE IF NOT EXISTS features (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		place_id INTEGER NOT NULL, 
		space INTEGER NOT NULL, 
		bathrooms INTEGER NOT NULL, 
		bedrooms INTEGER NOT NULL, 
		kitchen INTEGER NOT NULL, 
		capacity INTEGER NOT NULL,
		FOREIGN KEY (place_id) REFERENCES places (id) 
	)`
	_, err = DB.Exec(createFeaturesTable)

	if err != nil {
		panic(err)
	}

	createUsersTable := `CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT, 
		first_name TEXT NOT NULL,
		last_name TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	)`

	_, err = DB.Exec(createUsersTable)

	if err != nil {
		panic(err)
	}

	createOwnersTable := `CREATE TABLE IF NOT EXISTS owners (
		id INTEGER PRIMARY KEY AUTOINCREMENT, 
		first_name TEXT NOT NULL,
		last_name TEXT NOT NULL,
		phone TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	)`
	_, err = DB.Exec(createOwnersTable)

	if err != nil {
		panic(err)
	}

}
