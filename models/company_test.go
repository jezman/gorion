package models

import (
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
)

func TestCompany(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a strub database connection", err)
	}
	defer db.Close()

	app := &DB{db}
	rows := sqlmock.NewRows([]string{"Company"}).
		AddRow("company 1").
		AddRow("company 2")

	mock.ExpectQuery("SELECT Name FROM pCompany").
		WillReturnRows(rows)

	if _, err = app.Company(); err != nil {
		t.Errorf("error was not expected while gets company ", err)
	}
}
