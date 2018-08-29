package models

// Company model
type Company struct {
	Name string
	CountOfEmployees uint
}

// Company get all comanies from database
// return pionter to Company struct and error
func (db *DB) Company() ([]*Company, error) {
	query := `SELECT c.Name, Count(pList.Name) FROM pList
				JOIN pCompany c ON (c.ID = Company)
				GROUP BY c.Name`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var companies = make([]*Company, 0)
	for rows.Next() {
		company := new(Company)
		if err = rows.Scan(&company.Name, &company.CountOfEmployees); err != nil {
			return nil, err
		}

		companies = append(companies, company)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return companies, nil
}
