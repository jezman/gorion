package helpers

const (
	TestQueryEmployees = `^SELECT (.+), (.+), (.+), (.+) FROM pList p
		JOIN pCompany c ON ((.+) = (.+))
		ORDER BY (.+)$`
	TestQueryEmployeesByCompany = `^SELECT (.+), (.+), (.+), (.+) FROM pList
		JOIN pCompany c ON ((.+) = (.+))
		WHERE c.Name = (.+)
		ORDER BY (.+)$`
	TestQueryCompanies = `^SELECT (.+), (.+) FROM pList
		JOIN pCompany c ON ((.+) = (.+))
		GROUP BY (.+)$`
	TestQueryDoors = "^SELECT (.+), (.+) FROM AcessPoint ORDER BY GIndex$"
)
