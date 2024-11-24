package models

import (
	"time"

	"github.com/joshkiss/polyloggerclone/db"
)

type Entry struct {
	ID        int64
	Title     string
	Content   string
	Lang      string
	DateTime  time.Time
	TimeSpent int64
	Type      string
	UserId    int64
}

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
	return nil
}

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
