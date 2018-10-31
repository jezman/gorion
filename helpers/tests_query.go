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
	TestQueryEvents = `^SELECT (.+), (.+), (.+), (.+), (.+), (.+), (.+)
	    FROM pLogData l
		JOIN pList p ON ((.+) = (.+))
		JOIN pCompany c ON ((.+) = (.+))
		JOIN Events e ON ((.+) = (.+))
		JOIN AcessPoint a ON ((.+) = (.+))
		WHERE TimeVal BETWEEN (.+) AND (.+)
		AND (.+) BETWEEN 26 AND 29
		ORDER BY (.+)$`
	TestQueryEventsByEmployeeAndDoor = `^SELECT (.+), (.+), (.+), (.+), (.+), (.+), (.+)
		FROM pLogData l
		JOIN pList p ON ((.+) = (.+))
		JOIN pCompany c ON ((.+) = (.+))
		JOIN Events e ON ((.+) = (.+))
		JOIN AcessPoint a ON ((.+) = (.+))
		WHERE (.+) BETWEEN (.+) AND (.+)
		AND e.Event BETWEEN 26 AND 29
		AND p.Name = (.+)
		AND DoorIndex = (.+)
		ORDER BY (.+)$`
	TestQueryEventsByEmployee = `^SELECT (.+), (.+), (.+), (.+), (.+), (.+), (.+)
		FROM pLogData l
		JOIN pList p ON ((.+) = (.+))
		JOIN pCompany c ON ((.+) = (.+))
		JOIN Events e ON ((.+) = (.+))
		JOIN AcessPoint a ON ((.+) = (.+))
		WHERE (.+) BETWEEN (.+) AND (.+)
		AND e.Event BETWEEN 26 AND 29
		AND p.Name = (.+)
		ORDER BY (.+)$`
	TestQueryEventsByDoor = `^SELECT (.+), (.+), (.+), (.+), (.+), (.+), (.+)
		FROM pLogData l
		JOIN pList p ON ((.+) = (.+))
		JOIN pCompany c ON ((.+) = (.+))
		JOIN Events e ON ((.+) = (.+))
		JOIN AcessPoint a ON ((.+) = (.+))
		WHERE TimeVal BETWEEN (.+) AND (.+)
		AND e.Event BETWEEN 26 AND 29
		AND DoorIndex = (.+)
		ORDER BY (.+)$`
	TestQueryEventsDenied = `^SELECT (.+), (.+), (.+), (.+), (.+), (.+), (.+)
	    FROM pLogData l
		JOIN pList p ON ((.+) = (.+))
		JOIN pCompany c ON ((.+) = (.+))
		JOIN Events e ON ((.+) = (.+))
		JOIN AcessPoint a ON ((.+) = (.+))
		WHERE TimeVal BETWEEN (.+) AND (.+)
		AND (.+) IN (.+)
		ORDER BY (.+)$`
	TestQueryWorkedTime = `^SELECT (.+), (.+), (.+), (.+), (.+), (.+)
		FROM pLogData l
		JOIN pList p ON ((.+) = (.+))
		JOIN pCompany c ON ((.+) = (.+))
		WHERE TimeVal BETWEEN (.+) AND (.+)
		GROUP BY (.+), (.+), (.+), (.+), (.+)$`
	TestQueryWorkedTimeByCompany = `^SELECT (.+), (.+), (.+), (.+), (.+), (.+)
		FROM pLogData l
		JOIN pList p ON ((.+) = (.+))
		JOIN pCompany c ON ((.+) = (.+))
		WHERE TimeVal BETWEEN (.+) AND (.+)
		AND (.+) = (.+)
		GROUP BY (.+), (.+), (.+), (.+), (.+)$`
	TestQueryWorkedTimeByEmployee = `^SELECT (.+), (.+), (.+), (.+), (.+), (.+)
		FROM pLogData l
		JOIN pList p ON ((.+) = (.+))
		JOIN pCompany c ON ((.+) = (.+))
		WHERE TimeVal BETWEEN (.+) AND (.+)
		AND (.+) = (.+)
		GROUP BY (.+), (.+), (.+), (.+), (.+)$`
)
