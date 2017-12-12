package models

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
func (db *DB) Employees() ([]*Employee, error) {
	rows, err := db.Query(`SELECT p.Name, p.FirstName, p.MidName, c.Name FROM pList p
		JOIN pCompany c ON (c.ID = p.Company)
		ORDER BY c.Name`)
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
