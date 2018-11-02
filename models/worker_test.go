package models

import (
	"testing"

	"github.com/jezman/gorion/helpers"

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
		AddRow("f1", "m1", "l1", "c1").
		AddRow("f2", "m2", "l2", "c2")

	mock.ExpectQuery(helpers.TestQueryEmployees).WillReturnRows(rows)

	if _, err = app.Workers(""); err != nil {
		t.Errorf("error was not expected while gets worker %q ", err)
	}

	mock.ExpectQuery(helpers.TestQueryEmployeesByCompany).WithArgs("company").WillReturnRows(rows)

	if _, err = app.Workers("company"); err != nil {
		t.Errorf("error was not expected while gets worker by company %q ", err)
	}
}
