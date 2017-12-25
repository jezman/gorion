package models

import (
	"database/sql"
)

// Datastore methods
type Datastore interface {
	Company(string) ([]*Company, error)
	Doors(string) ([]*Door, error)
	Employees(string) ([]*Employee, error)
	Events(string) ([]*Event, error)
	WorkedTime(string) ([]*Event, error)
}

// DB structure used as receiver in methods
type DB struct {
	*sql.DB
}

// OpenDB opening connecting to server
// return pointer to struct DB and error
func OpenDB(dsn string) (*DB, error) {
	db, err := sql.Open("mssql", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return &DB{db}, nil
}
