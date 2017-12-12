package models

import (
	"database/sql"
)

// Datastore methods
type Datastore interface {
	// Company() ([]*Company, error)
	// Doors() ([]*Door, error)
	// Employees() ([]*Employee, error)
	// Events(string) ([]*Event, error)
	// TimeWorked(string) ([]*Event, error)
}

// Config file structure
type Config struct {
	Server   string `json:"server"`
	Database string `json:"database"`
	User     string `json:"user"`
	Password string `json:"password"`
}

// DB structure used as receiver in methods
type DB struct {
	*sql.DB
}
