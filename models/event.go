package models

import (
	"time"
)

// Event model
type Event struct {
	Employee   Employee
	FirstTime  time.Time
	LastTime   time.Time
	Company    Company
	Door       Door
	Action     string
	WorkedTime time.Duration
}

// Events gets the list of events for the time period
// return pionter to Event struct and error
func (db *DB) Events(query string) ([]*Event, error) {
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events = make([]*Event, 0)
	for rows.Next() {
		event := new(Event)
		err := rows.Scan(
			&event.Employee.LastName,
			&event.Employee.FirstName,
			&event.Employee.MidName,
			&event.Company.Name,
			&event.FirstTime,
			&event.Door.Name,
			&event.Action,
		)

		if err != nil {
			return nil, err
		}
		event.Employee.FullName = event.Employee.LastName + " " +
			event.Employee.FirstName + " " + event.Employee.MidName

		event.WorkedTime = event.LastTime.Sub(event.FirstTime)

		events = append(events, event)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return events, nil

}
