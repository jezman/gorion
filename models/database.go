package models

import (
	"database/sql"
	"encoding/json"
	"os"
)

// Datastore methods
type Datastore interface {
	Company() ([]*Company, error)
	Doors() ([]*Door, error)
	Employees() ([]*Employee, error)
	Events(string) ([]*Event, error)
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

// Read open file and parsing JSON
// return data source name
func (c *Config) Read(file string) string {
	confFile, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer confFile.Close()

	decoder := json.NewDecoder(confFile)
	err = decoder.Decode(&c)
	if err != nil {
		panic(err)
	}

	dsn := "server=" + c.Server + ";user id=" + c.User +
		";password=" + c.Password + ";database=" + c.Database

	return dsn
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
