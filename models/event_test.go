package models

import (
	"testing"
	"time"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
)

func TestEvents(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a strub database connection", err)
	}
	defer db.Close()

	app := &DB{DB: db}

	column := []string{"Time", "firstName", "midName", "lastName", "Company", "Event", "Door"}
	rows := sqlmock.NewRows(column).
		AddRow(
			"firstName",
			"midName",
			"lastName",
			"company",
			time.Now(),
			"action",
			"door",
		)

	query = "SELECT p.Name, p.FirstName, p.MidName, c.Name, TimeVal, e.Contents, a.Name FROM pLogData l JOIN pList p ON (p.ID = l.HozOrgan) JOIN pCompany c ON (c.ID = p.Company) JOIN Events e ON (e.Event = l.Event) JOIN AcessPoint a ON (a.GIndex = l.DoorIndex) WHERE TimeVal BETWEEN \\? AND \\? AND e.Event BETWEEN 26 AND 29 ORDER BY TimeVal"

	mock.ExpectQuery(query).
		WithArgs("22.08.2018", "23.08.2018").
		WillReturnRows(rows)

	if _, err = app.Events("22.02.2018", "23.03.2018", "", 0); err != nil {
		t.Errorf("error was not expected while gets events %q ", err)
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
		AddRow(
			"firstName",
			"midName",
			"lastName",
			"company",
			time.Now(),
			time.Now(),
		)
	query := "SELECT p.Name, p.FirstName, p.MidName, c.Name, min(TimeVal), max(TimeVal) FROM pLogData l JOIN pList p ON (p.ID = l.HozOrgan) JOIN pCompany c ON (c.ID = p.Company) WHERE TimeVal BETWEEN ? AND ? GROUP BY p.Name, p.FirstName, p.MidName, c.Name, CONVERT(varchar(20), TimeVal, 104)"

	mock.ExpectQuery(query).WillReturnRows(rows)

	if _, err = app.WorkedTime("22.02.2018", "23.03.2018", ""); err != nil {
		t.Errorf("error was not expected while gets worked time %q ", err)
	}

}
