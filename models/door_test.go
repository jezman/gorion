package models

import (
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
)

func TestDoors(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a strub database connection", err)
	}
	defer db.Close()

	app := &DB{db}
	rows := sqlmock.NewRows([]string{"GIndxex", "Name"}).
		AddRow(1, "door 1")

	query := "SELECT GIndex, Name FROM AcessPoint"
	mock.ExpectQuery(query).
		WillReturnRows(rows)

	if _, err = app.Doors(query); err != nil {
		t.Errorf("error was not expected while gets doors %q ", err)
	}
}
