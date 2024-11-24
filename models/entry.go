package models

import "time"

type Entry struct {
	ID        int64
	Title     string
	Content   string
	Lang      string
	DateTime  time.Time
	TimeSpent int64
	Type      string
}

func (e *Entry) Save() error {

	return nil
}

func (e *Entry) Update() error {
	return nil
}

func GetAllEntries() ([]Entry, error) {

	return nil, nil
}

func GetAllByUser(userId int64) {}
