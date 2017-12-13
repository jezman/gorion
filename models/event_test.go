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

	app := &DB{db}

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
	query := "select TimeVal, HozOrgan, Remark, DoorIndex, NumCom, Event, indexKey from pLogData"

	mock.ExpectQuery(query).WillReturnRows(rows)

	if _, err = app.Events(query); err != nil {
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

	app := &DB{db}

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
	query := "select TimeVal, HozOrgan, Remark, DoorIndex, NumCom, Event from pLogData"

	mock.ExpectQuery(query).WillReturnRows(rows)

	if _, err = app.WorkedTime(query); err != nil {
		t.Errorf("error was not expected while gets worked time %q ", err)
	}

}
