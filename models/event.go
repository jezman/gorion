package models

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/jezman/gorion/helpers"
)

var (
	rows  *sql.Rows
	err   error
	query string
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
// return pointer to Event struct and error
func (db *DB) Events(firstDate, lastDate, employee string, door uint, denied bool) ([]*Event, error) {
	// change the query depending on the input flag
	switch {
	case door != 0 && employee != "":
		if !helpers.ValidationEmployee(employee) {
			fmt.Print("invalid employee. allowed only letters")
			os.Exit(1)
		}
		rows, err = db.Query(helpers.QueryEventsBeEmployeeAndDoor, firstDate, lastDate, employee, door)
	case employee != "":
		if !helpers.ValidationEmployee(employee) {
			fmt.Print("invalid employee. allowed only letters")
			os.Exit(1)
		}
		rows, err = db.Query(helpers.QueryEventsByEmployee, firstDate, lastDate, employee)
	case door != 0:
		rows, err = db.Query(helpers.QueryEventsByDoor, firstDate, lastDate, door)
	case denied:
		rows, err = db.Query(helpers.QueryEventsDenied, firstDate, lastDate)
	default:
		rows, err = db.Query(helpers.QueryEvents, firstDate, lastDate)
	}

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events = make([]*Event, 0)
	for rows.Next() {
		event := new(Event)
		err = rows.Scan(
			&event.Employee.LastName,
			&event.Employee.FirstName,
			&event.Employee.MidName,
			&event.Company.Name,
			&event.FirstTime,
			&event.Action,
			&event.Door.Name,
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

// WorkedTime gets the list of employees and
// calculates their worked time
// return pointer to Event struct and error
func (db *DB) WorkedTime(firstDate, lastDate, employee, company string) ([]*Event, error) {
	if !helpers.ValidationDate(firstDate) || !helpers.ValidationDate(lastDate) {
		fmt.Print("invalid date. corrects format: DD.MM.YYYY or DD-MM-YYYY")
		os.Exit(1)
	}

	switch {
	case employee != "":
		if !helpers.ValidationEmployee(employee) {
			fmt.Print("invalid employee. allowed only letters")
			os.Exit(1)
		}
		rows, err = db.Query(helpers.QueryWorkedTimeByEmployee, firstDate, lastDate, employee)
	case company != "":
		rows, err = db.Query(helpers.QueryWorkedTimeByCompany, firstDate, lastDate, company)
	default:
		rows, err = db.Query(helpers.QueryWorkedTime, firstDate, lastDate)
	}

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events = make([]*Event, 0)
	for rows.Next() {
		event := new(Event)
		err = rows.Scan(
			&event.Employee.LastName,
			&event.Employee.FirstName,
			&event.Employee.MidName,
			&event.Employee.Company.Name,
			&event.FirstTime,
			&event.LastTime,
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
