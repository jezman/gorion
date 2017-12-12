package models

// Company model
type Company struct {
	Name string
}

// Company get all comanies from database
// return pionter to Company struct and error
func (db *DB) Company() ([]*Company, error) {
	rows, err := db.Query("SELECT Name FROM pCompany ORDER BY Name")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var company = make([]*Company, 0)
	for rows.Next() {
		c := new(Company)
		if err = rows.Scan(&c.Name); err != nil {
			return nil, err
		}

		company = append(company, c)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return company, nil
}
