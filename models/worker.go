package models

import (
	"database/sql"

	"github.com/jezman/gorion/helpers"
)

// Worker model
type Worker struct {
	FirstName string
	LastName  string
	MidName   string
	FullName  string
	Company   Company
}

// Workers get all workers from database
// return pionter to Worker struct and error
func (db *DB) Workers(companyName string) ([]*Worker, error) {
	var (
		rows *sql.Rows
		err  error
	)

	if companyName != "" {
		rows, err = db.Query(helpers.QueryEmployeesByCompany, companyName)
	} else {
		rows, err = db.Query(helpers.QueryEmployees)
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var workers = make([]*Worker, 0)
	for rows.Next() {
		w := new(Worker)
		if err = rows.Scan(&w.LastName, &w.FirstName, &w.MidName, &w.Company.Name); err != nil {
			return nil, err
		}

		w.FullName = w.LastName + " " + w.FirstName + " " + w.MidName

		workers = append(workers, w)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return workers, nil
}
