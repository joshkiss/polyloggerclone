package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./api.db")

	if err != nil {
		panic("Could not connect to database.")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	// write SQL to make tables here
	// users
	// entries

	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT NOT NULL UNIQUE,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	)
	`

	createEntriesTable := `
	CREATE TABLE IF NOT EXISTS entries (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	title TEXT NOT NULL, 
	content TEXT NOT NULL, 
	lang TEXT NOT NULL, 
	datetime DATETIME NOT NULL, 
	timespent INTEGER NOT NULL, 
	type TEXT NOT NULL,
	user_id INTEGER, FOREIGN KEY (user_id) REFERENCES user(id)
	)
	`

	_, err := DB.Exec(createUsersTable)
	if err != nil {
		panic("Could not create 'users' table.")
	}
	_, err = DB.Exec(createEntriesTable)
	if err != nil {
		panic("Could not create 'users' table.")
	}
}
