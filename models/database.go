package models

import (
	"database/sql"
)

// Datastore methods
type Datastore interface {
	Company() ([]*Company, error)
	Doors() ([]*Door, error)
	Employees(string) ([]*Employee, error)
	Events(string, string, string, uint, bool) ([]*Event, error)
	WorkedTime(string, string, string, string) ([]*Event, error)
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

	return &DB{DB: db}, nil
}
