package models

import (
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
)

func TestEmployees(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a strub database connection", err)
	}
	defer db.Close()

	app := &DB{db}

	column := []string{"firstName", "midName", "lastName", "Company"}
	rows := sqlmock.NewRows(column).
		AddRow(
			"firstName",
			"midName",
			"lastName",
			"company",
		)
	query := "SELECT Name, FirstName, MidName, Company FROM pList"
	mock.ExpectQuery(query).WillReturnRows(rows)

	if _, err = app.Employees(query); err != nil {
		t.Errorf("error was not expected while gets events %q ", err)
	}
}
