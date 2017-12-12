package query

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/jezman/gorion/check"
)

// Events check flags and return query for events list
func Events(doorID uint, employee, firstDate, lastDate string) string {
	pName := " AND p.Name = '"
	doorIndex := "' AND DoorIndex = "
	orderBy := "' ORDER BY TimeVal"

	query := []string{
		"SELECT p.Name, p.FirstName, p.MidName, c.Name, TimeVal, e.Contents, a.Name ",
		"FROM pLogData l ",
		"JOIN pList p ON (p.ID = l.HozOrgan) ",
		"JOIN pCompany c ON (c.ID = p.Company) ",
		"JOIN Events e ON (e.Event = l.Event) ",
		"JOIN AcessPoint a ON (a.GIndex = l.DoorIndex) ",
		"WHERE TimeVal BETWEEN '", firstDate, "' AND '", lastDate, "'",
		" AND e.Event BETWEEN 26 AND 29",
		"ORDER BY TimeVal",
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

	// query changer
	add := func(cmd ...string) {
		query = append(query[:len(query)-1], cmd...)
	}

	// change the query depending on the input flag
	if doorID != 0 && employee != "" {
		// check employee flag
		if err := check.Employee(employee); err != nil {
			fmt.Printf("'%s' %s\n", employee, err)
			os.Exit(1)
		}
		add(pName, employee, doorIndex, strconv.Itoa(int(doorID)), orderBy[1:])
	} else if employee != "" {
		add(pName, employee, orderBy)
	} else if doorID != 0 {
		add(doorIndex[1:], strconv.Itoa(int(doorID)), orderBy[1:])
	}

	return strings.Join(query, "")
}
