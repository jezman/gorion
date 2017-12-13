package models

import (
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
)

func TestCompany(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a strub database connection", err)
	}
	defer db.Close()

	app := &DB{db}
	rows := sqlmock.NewRows([]string{"Company"}).
		AddRow("company 1")

	query := "SELECT Name FROM pCompany"
	mock.ExpectQuery(query).
		WillReturnRows(rows)

	if _, err = app.Company(query); err != nil {
		t.Errorf("error was not expected while gets company %q ", err)
	}
}
