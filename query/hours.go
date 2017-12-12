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
	if err := check.Date(firstDate); err != nil {
		fmt.Printf("'%s' %s\n", firstDate, err)
		os.Exit(1)
	}

	if err := check.Date(lastDate); err != nil {
		fmt.Printf("'%s' %s\n", lastDate, err)
		os.Exit(1)
	}

	// change query if emploeey is empty
	if employee == "" {
		// remove  "AND p.Name = 'employee' from query"
		query = append(query[:9], query[12])
	} else if err := check.Employee(employee); err != nil {
		fmt.Printf("'%s' %s\n", employee, err)
		os.Exit(1)
	}

	return strings.Join(query, "")
}
