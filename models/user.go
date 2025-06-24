package models

import (
	"errors"

	"github.com/damianpugliese/go_rest_api_events_booking/db"
	"github.com/damianpugliese/go_rest_api_events_booking/utils"
)

type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *User) Save() error {
	query := `
		INSERT INTO users (email, password)
		VALUES (?, ?)
	`
	
	stmt, err := db.DB.Prepare(query)
	
	if err != nil {
		return err
	}
	
	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)

	if err != nil {	
		return err
	}

	result, err := stmt.Exec(u.Email, hashedPassword)
	
	if err != nil {
		return err
	}
	
	id, err := result.LastInsertId()
	u.ID = id

	return err
}

func (u *User) ValidateCredentials() error {
	query := `
		SELECT password FROM users 
		WHERE email = ?
	`
	
	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string

	err := row.Scan(&retrievedPassword)
	
	if err != nil {
		return errors.New("invalid credentials")
	}
	
	passwordMatch := utils.CheckPasswordHash(u.Password, retrievedPassword)

	if !passwordMatch {
		return errors.New("invalid credentials")
	}

	return nil
}