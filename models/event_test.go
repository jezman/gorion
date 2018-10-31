package models

import (
	"testing"
	"time"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/jezman/gorion/helpers"
)

var (
	timeNow = time.Now().Local()
	firstDate = timeNow.Format("02.01.2006")
	lastDate = timeNow.AddDate(0, 0, 1).Format("02.01.2006")
	employee = "Employee"
	company = "Company"
	door = uint(22)
	denied = true
)

func TestEvents(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a strub database connection", err)
	}
	defer db.Close()

	app := &DB{DB: db}


	column := []string{"Time", "firstName", "midName", "lastName", "Company", "Door", "Event"}
	rows := sqlmock.NewRows(column).
		AddRow("firstName", "midName", "lastName", "company", time.Now(), "door", "action")

	mock.ExpectQuery(helpers.TestQueryEvents).
		WithArgs(firstDate, lastDate).
		WillReturnRows(rows)

	if _, err := app.Events(firstDate, lastDate, "", 0, false); err != nil {
		t.Errorf("error was not expected while gets all events %q ", err)
	}

	mock.ExpectQuery(helpers.TestQueryEventsByEmployeeAndDoor).
		WithArgs(firstDate, lastDate, employee, door).
		WillReturnRows(rows)

	if _, err = app.Events(firstDate, lastDate, employee, door, false); err != nil {
		t.Errorf("error was not expected while gets events by employee and door %q ", err)
	}

	mock.ExpectQuery(helpers.TestQueryEventsByEmployee).
		WithArgs(firstDate, lastDate, employee).
		WillReturnRows(rows)

	if _, err = app.Events(firstDate, lastDate, employee, 0, false); err != nil {
		t.Errorf("error was not expected while gets events by employee %q ", err)
	}

	mock.ExpectQuery(helpers.TestQueryEventsByDoor).
		WithArgs(firstDate, lastDate, door).
		WillReturnRows(rows)

	if _, err = app.Events(firstDate, lastDate, "", door, false); err != nil {
		t.Errorf("error was not expected while gets events by door %q ", err)
	}

	mock.ExpectQuery(helpers.TestQueryEventsDenied).
		WithArgs(firstDate, lastDate).
		WillReturnRows(rows)

	if _, err = app.Events(firstDate, lastDate, "", 0, denied); err != nil {
		t.Errorf("error was not expected while gets denied events %q ", err)
	}
}

func TestWorkedTime(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a strub database connection", err)
	}
	defer db.Close()

	app := &DB{DB: db}

	column := []string{"Time", "firstName", "midName", "lastName", "Company", "Event"}
	rows := sqlmock.NewRows(column).
		AddRow("firstName",	"midName", "lastName", "company", timeNow, timeNow)

	mock.ExpectQuery(helpers.TestQueryWorkedTime).WillReturnRows(rows)

	if _, err = app.WorkedTime(firstDate, lastDate, "", ""); err != nil {
		t.Errorf("error was not expected while gets worked time %q ", err)
	}

	mock.ExpectQuery(helpers.TestQueryWorkedTimeByCompany).WillReturnRows(rows)

	if _, err = app.WorkedTime(firstDate, lastDate, "", company); err != nil {
		t.Errorf("error was not expected while gets worked time %q ", err)
	}

	mock.ExpectQuery(helpers.TestQueryWorkedTimeByEmployee).WillReturnRows(rows)

	if _, err = app.WorkedTime(firstDate, lastDate, employee, ""); err != nil {
		t.Errorf("error was not expected while gets worked time %q ", err)
	}
}
