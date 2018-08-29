package models

import (
	"database/sql"
)

// Employee model
type Employee struct {
	FirstName string
	LastName  string
	MidName   string
	FullName  string
	Company   Company
}

// Employees get all employees from database
// return pionter to Employee struct and error
func (db *DB) Employees(companyName string) ([]*Employee, error) {
	var (
		rows  *sql.Rows
		err   error
		query string
	)

	if companyName != "" {
		query = `SELECT plist.Name, pList.FirstName, pList.MidName, c.Name from pList
						JOIN pCompany c ON (c.ID = Company)
						WHERE c.Name = ?
						ORDER BY pList.Name`
		rows, err = db.Query(query, companyName)

	} else {
		query = `SELECT p.Name, p.FirstName, p.MidName, c.Name FROM pList p
					JOIN pCompany c ON (c.ID = p.Company)
					ORDER BY c.Name`
		rows, err = db.Query(query)
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var employees = make([]*Employee, 0)
	for rows.Next() {
		e := new(Employee)
		if err = rows.Scan(&e.LastName, &e.FirstName, &e.MidName, &e.Company.Name); err != nil {
			return nil, err
		}

		e.FullName = e.LastName + " " + e.FirstName + " " + e.MidName

		employees = append(employees, e)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return employees, nil
}
