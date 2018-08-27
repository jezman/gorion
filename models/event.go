package models

import (
	"time"
	"fmt"
	"os"
	"github.com/jezman/gorion/check"
	"database/sql"
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
func (db *DB) Events(firstDate, lastDate, employee string, door uint) ([]*Event, error) {
	// change the query depending on the input flag
	switch {
	case door != 0 && employee != "":
		// check employee flag
		if !check.Employee(employee) {
			fmt.Print("invalid employee. allowed only letters")
			os.Exit(1)
		}

		query = `SELECT p.Name, p.FirstName, p.MidName, c.Name, TimeVal, e.Contents, a.Name
				FROM pLogData l
				JOIN pList p ON (p.ID = l.HozOrgan)
				JOIN pCompany c ON (c.ID = p.Company)
				JOIN Events e ON (e.Event = l.Event)
				JOIN AcessPoint a ON (a.GIndex = l.DoorIndex)
				WHERE TimeVal BETWEEN ? AND ?
				AND e.Event BETWEEN 26 AND 29
				AND p.Name = ?
				AND DoorIndex = ?
				ORDER BY TimeVal`
		rows, err = db.Query(query, firstDate, lastDate, employee, door)

	case employee != "":
		if !check.Employee(employee) {
			fmt.Print("invalid employee. allowed only letters")
			os.Exit(1)
		}
		// add employee to query
		query = `SELECT p.Name, p.FirstName, p.MidName, c.Name, TimeVal, e.Contents, a.Name
				FROM pLogData l
				JOIN pList p ON (p.ID = l.HozOrgan)
				JOIN pCompany c ON (c.ID = p.Company)
				JOIN Events e ON (e.Event = l.Event)
				JOIN AcessPoint a ON (a.GIndex = l.DoorIndex)
				WHERE TimeVal BETWEEN ? AND ?
				AND e.Event BETWEEN 26 AND 29
				AND p.Name = ?
				ORDER BY TimeVal`

		rows, err = db.Query(query, firstDate, lastDate, employee)

	case door != 0:
		// add door to query
		query = `SELECT p.Name, p.FirstName, p.MidName, c.Name, TimeVal, e.Contents, a.Name
				FROM pLogData l
				JOIN pList p ON (p.ID = l.HozOrgan)
				JOIN pCompany c ON (c.ID = p.Company)
				JOIN Events e ON (e.Event = l.Event)
				JOIN AcessPoint a ON (a.GIndex = l.DoorIndex)
				WHERE TimeVal BETWEEN ? AND ?
				AND e.Event BETWEEN 26 AND 29
				AND DoorIndex = ?
				ORDER BY TimeVal`
		rows, err = db.Query(query, firstDate, lastDate, door)

	default:
		query = `SELECT p.Name, p.FirstName, p.MidName, c.Name, TimeVal, e.Contents, a.Name
				FROM pLogData l
				JOIN pList p ON (p.ID = l.HozOrgan)
				JOIN pCompany c ON (c.ID = p.Company)
				JOIN Events e ON (e.Event = l.Event)
				JOIN AcessPoint a ON (a.GIndex = l.DoorIndex)
				WHERE TimeVal BETWEEN ? AND ?
				AND e.Event BETWEEN 26 AND 29
				ORDER BY TimeVal`
		rows, err = db.Query(query, firstDate, lastDate)
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

// WorkedTime gets the list of employees and
// calculates their worked time
// return pointer to Event struct and error
func (db *DB) WorkedTime(firstDate, lastDate, employee string) ([]*Event, error) {
	// check dates
	if !check.Date(firstDate) || !check.Date(lastDate) {
		fmt.Print("invalid date. corrects format: DD.MM.YYYY or DD-MM-YYYY")
		os.Exit(1)
	}

	// change query if employee is not empty
	if employee != "" {
		if !check.Employee(employee) {
			fmt.Print("invalid employee. allowed only letters")
			os.Exit(1)
		}

		query = `SELECT p.Name, p.FirstName, p.MidName, c.Name, min(TimeVal), max(TimeVal)
			FROM pLogData l
			JOIN pList p ON (p.ID = l.HozOrgan)
			JOIN pCompany c ON (c.ID = p.Company)
			WHERE TimeVal BETWEEN ? AND ?
			AND p.Name = ?
			GROUP BY p.Name, p.FirstName, p.MidName, c.Name, CONVERT(varchar(20), TimeVal, 104)`

		rows, err = db.Query(query, firstDate, lastDate, employee)
	} else {
		query = `SELECT p.Name, p.FirstName, p.MidName, c.Name, min(TimeVal), max(TimeVal)
			FROM pLogData l
			JOIN pList p ON (p.ID = l.HozOrgan)
			JOIN pCompany c ON (c.ID = p.Company)
			WHERE TimeVal BETWEEN ? AND ?
			GROUP BY p.Name, p.FirstName, p.MidName, c.Name, CONVERT(varchar(20), TimeVal, 104)`

		rows, err = db.Query(query, firstDate, lastDate)
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
