package query

import (
	"fmt"
	"os"
	"strings"

	"github.com/jezman/gorion/check"
)

// WorkedTime check flags and return query for worked time list
func WorkedTime(employee, firstDate, lastDate string) string {
	query := []string{
		"SELECT p.Name, p.FirstName, p.MidName, c.Name, min(TimeVal), max(TimeVal) ",
		"FROM pLogData l ",
		"JOIN pList p ON (p.ID = l.HozOrgan) ",
		"JOIN pCompany c ON (c.ID = p.Company) ",
		"WHERE TimeVal BETWEEN '", firstDate, "' AND '", lastDate, "'",
		" AND p.Name = '", employee, "'",
		" GROUP BY p.Name, p.FirstName, p.MidName, c.Name, CONVERT(varchar(20), TimeVal, 104)",
	}

	// check dates
	if !check.Date(firstDate) || !check.Date(lastDate) {
		fmt.Print("invalid date. corrects format: DD.MM.YYYY or DD-MM-YYYY")
		os.Exit(1)
	}

	// change query if emploeey is empty
	if employee == "" {
		// remove  "AND p.Name = 'employee' from query"
		query = append(query[:9], query[12])

		// check employee
	} else if !check.Employee(employee) {
		fmt.Print("invalid employee. allowed only latters")
		os.Exit(1)
	}

	return strings.Join(query, "")
}
