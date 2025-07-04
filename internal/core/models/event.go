package models

import (
	"time"

	"github.com/damianpugliese/go_rest_api_events_booking/internal/infrastructure/db"
)

type Event struct {
	ID          int64 	 	`json:"id"`
	Name        string 		`binding:"required" json:"name"`
	Description string 		`binding:"required" json:"description"`
	Location    string 		`binding:"required" json:"location"`
	DateTime    time.Time `binding:"required" json:"datetime"`
	UserID      int64     `json:"user_id"`
}

func (e *Event) Save() error {
	query := `
		INSERT INTO events (name, description, location, datetime, user_id)
		VALUES (?, ?, ?, ?, ?)
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	e.ID = id

	return err
}

func GetAllEvents() ([]Event, error) {
	query := `
		SELECT * FROM events
	`

	rows, err := db.DB.Query(query)
	
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	events := []Event{}

	for rows.Next() {
		var event Event
		err = rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}

func GetEventById(id int64) (*Event, error) {
	query := `
		SELECT * FROM events WHERE id = ?
	`
	
	row := db.DB.QueryRow(query, id)
	
	var event Event

	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
	
	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (e Event) Update() error {
	query := `
		UPDATE events 
		SET name = ?, description = ?, location = ?, datetime = ? 
		WHERE id = ?
	`
	
	stmt, err := db.DB.Prepare(query)
	
	if err != nil {
		return err
	}
	
	defer stmt.Close()

	_, err = stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.ID)
	
	if err != nil {
		return err
	}

	return err
}

func (e Event) Delete() error {
	query := `
		DELETE FROM events 
		WHERE id = ?
	`
	
	stmt, err := db.DB.Prepare(query)
	
	if err != nil {
		return err
	}
	
	defer stmt.Close()

	_, err = stmt.Exec(e.ID)
	
	if err != nil {
		return err
	}

	return err
}

func (e *Event) Register(userId int64) error {
	query := `
		INSERT INTO registrations (event_id, user_id)
		VALUES (?, ?)
	`
	
	stmt, err := db.DB.Prepare(query)
	
	if err != nil {
		return err
	}
	
	defer stmt.Close()
	
	result, err := stmt.Exec(e.ID, userId)
	
	if err != nil {
		return err
	}
	
	id, err := result.LastInsertId()
	e.ID = id

	return err
}

func (e *Event) Cancel(userId int64) error {
	query := `
		DELETE FROM registrations
		WHERE event_id = ? AND user_id = ?
	`
	
	stmt, err := db.DB.Prepare(query)
	
	if err != nil {
		return err
	}
	
	defer stmt.Close()
	
	_, err = stmt.Exec(e.ID, userId)
	
	return err
}
