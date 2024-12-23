package models

import (
	"time"

	"github.com/joshkiss/polyloggerclone/db"
)

type Entry struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Lang      string    `json:"lang"`
	DateTime  time.Time `json:"datetime"`
	TimeSpent int64     `json:"time_spent"`
	Type      string    `json:"type"`
	UserId    int64     `json:"user_id"`
}

// CRUD Functionality

func (e *Entry) Save() error {
	query := `INSERT INTO entries 
	(title, content, lang, datetime, timespent, type, user_id)
	VALUES (?, ?, ?, ?, ?, ?, ?)
	`
	statement, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer statement.Close()

	res, err := statement.Exec(e.Title, e.Content, e.Lang, e.DateTime, e.TimeSpent, e.Type, e.UserId)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	e.ID = id
	return nil
}

func (e *Entry) Update() error {
	query := `
	UPDATE entries 
	SET title = ?, content = ?, lang = ?, datetime = ?, timespent = ?, type = ?, user_id = ? 
	WHERE id = ?
	`
	statement, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer statement.Close()
	_, err = statement.Exec(e.Title, e.Content, e.Lang, e.DateTime, e.TimeSpent, e.Type, e.UserId, e.ID)
	return err
}

func (e Entry) Delete() error {
	query := "DELETE FROM entries WHERE id = ?"
	statement, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(e.ID)
	if err != nil {
		return err
	}
	return nil
}

// Get entries

func GetAllEntries() ([]Entry, error) {
	query := "SELECT * FROM entries"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var entries []Entry
	for rows.Next() {
		var entry Entry
		err = rows.Scan(&entry.ID, &entry.Title, &entry.Content, &entry.Lang, &entry.DateTime, &entry.TimeSpent, &entry.Type, &entry.UserId)
		if err != nil {
			return nil, err
		}

		entries = append(entries, entry)
	}

	return entries, nil
}

func GetAllByUser(userId int64) ([]Entry, error) {
	query := "SELECT * FROM entries WHERE user_id = ?"
	rows, err := db.DB.Query(query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var entries []Entry
	for rows.Next() {
		var entry Entry
		err = rows.Scan(&entry.ID, &entry.Title, &entry.Content, &entry.Lang, &entry.DateTime, &entry.TimeSpent, &entry.Type, &entry.UserId)
		if err != nil {
			return nil, err
		}

		entries = append(entries, entry)
	}

	return entries, nil
}

func GetEntryByID(id int64) (*Entry, error) {
	query := "SELECT * FROM entries WHERE id = ?"
	row := db.DB.QueryRow(query, id)

	var entry Entry
	err := row.Scan(&entry.ID, &entry.Title, &entry.Content, &entry.Lang, &entry.DateTime, &entry.TimeSpent, &entry.Type, &entry.UserId)
	if err != nil {
		return nil, err
	}

	return &entry, nil
}
