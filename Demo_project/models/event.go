package models

import (
	"crud_project/db"
	"time"
	_ "github.com/pelletier/go-toml/query"
)

// 
type Event struct{
	ID int64				
	Name string         `binding : "required"` 
	Description string  `binding : "required"`
	Location string     `binding : "required"`
	DateTime time.Time  `binding : "required"`
	UserID int
}

var events = []Event{} // empty slice to save events

func (e Event) Save() error{ 
	//Function to add the event data to the event slice

	query := `INSERT INTO events(name, description, location, dateTime, user_id)
	VALUES (?, ?, ?, ?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil{
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil{
		return err
	}
	id, err := result.LastInsertId()
	e.ID = id
	return err
	// events = append(events, e)
}

func GetAllEvents() ([]Event, error) {
	// Read all event from database.
	query:= "select * from events"
	rows, err := db.DB.Query(query)
	if err != nil{
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var event Event
		err:= rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		if err != nil{
			return nil, err
		}
		events= append(events, event)
	}

	return events, nil
}

func GetEventByID(id int64) (*Event, error){
	// Read event by ID.

	query := `select * from events where id = ?`
	row := db.DB.QueryRow(query, id) // for single row
	
	var event Event
	err:= row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
	if err != nil{
		return nil, err
	}
	
	return &event, nil

}

func (e Event) Delete() error{
	// Delete event from database.

	query := `DELETE FROM events WHERE id = ?`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(e.ID)
	return err
}

func (e Event) Update() (error){
	// Update an event.
	query := `
	UPDATE events 
	SET name = ?, description= ?, location = ?, dateTime = ?
	WHERE id = ?
	`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.ID)
	
	return err
}