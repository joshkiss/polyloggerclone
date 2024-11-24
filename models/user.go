package models

import (
	"errors"

	"github.com/joshkiss/polyloggerclone/db"
	"github.com/joshkiss/polyloggerclone/utils"
)

type User struct {
	ID       int64
	Username string `json:"username"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (u *User) Save() error {
	query := "INSERT INTO users (username, email, password) VALUES (?, ?, ?)"
	statement, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer statement.Close()

	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	res, err := statement.Exec(u.Username, u.Email, hashedPassword)
	if err != nil {
		return err
	}

	userId, err := res.LastInsertId()
	u.ID = userId
	return err
}

func (u *User) Validate() error {
	query := "SELECT id, password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email)

	var password string
	err := row.Scan(&u.ID, &password)
	if err != nil {
		return errors.New("invalid credentials")
	}

	passwordIsValid := utils.CheckPasswordHash(u.Password, password)
	if !passwordIsValid {
		return errors.New("invalid credentials")
	}

	return nil
}
