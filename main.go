package main

import (
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
)

type config struct {
	Server   string
	Database string
	User     string
	Password string
}

func main() {
	timeNow := time.Now().Local()
	dayFirstHour := timeNow.Format("02.01.2006") + " 00:00"
	dayLastHour := timeNow.Format("02.01.2006") + " 23:59"

	// Cli flags
	name := flag.String("name", "", "User last name.")
	door := flag.String("door", "", "List all doors with indexes.")
	doors := flag.Bool("list", false, "List all doors.")
	startDate := flag.String("start", dayFirstHour, "Start of time rande. Format: \"31.12.2017 23:59\"")
	endDate := flag.String("end", dayLastHour, "End of time range. Format: \"31.12.2017 23:59\"")

	flag.Parse()

	// Read config file.
	confFile, _ := os.Open("config.json")
	defer confFile.Close()

	decoder := json.NewDecoder(confFile)
	conf := config{}
	err := decoder.Decode(&conf)
	if err != nil {
		fmt.Println("Json decode error:", err)
	}

	// Connect to db
	dsn := "server=" + conf.Server + ";user id=" + conf.User + ";password=" + conf.Password + ";database=" + conf.Database
	db, err := sql.Open("mssql", dsn)
	if err != nil {
		fmt.Println("Cannot connect: ", err.Error())
		return
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("Cannot connect: ", err.Error())
		return
	}
	defer db.Close()

	query := []string{"SELECT p.Name AS Фамилия, p.FirstName AS Имя, P.MidName AS Отчество, TimeVal AS Время, e.Contents AS Событие, a.Name AS Дверь ",
		"FROM dbo.pLogData l ",
		"JOIN dbo.pList p ON (p.ID = l.HozOrgan) ",
		"JOIN dbo.Events e ON (e.Event = l.Event) ",
		"JOIN dbo.AcessPoint a ON (a.GIndex = l.DoorIndex) ",
		"WHERE TimeVal BETWEEN '", *startDate, "' AND '", *endDate, "'",
		" AND e.Event BETWEEN 26 AND 29",
		"ORDER BY TimeVal",
	}
	if *name != "" {
		query = append(query[0:len(query)-1], " AND p.Name = '", *name, "' ORDER BY TimeVal")
	}
	if *door != "" {
		query = append(query[0:len(query)-1], " AND DoorIndex = ", *door, " ORDER BY TimeVal")
	}
	if *doors != false {
		doorsList := "SELECT GIndex, Name from dbo.AcessPoint"
		execute(db, doorsList)
	}

	execute(db, strings.Join(query, ""))
}

func execute(db *sql.DB, cmd string) error {
	rows, err := db.Query(cmd)
	if err != nil {
		return err
	}
	defer rows.Close()
	cols, err := rows.Columns()
	if err != nil {
		return err
	}
	if cols == nil {
		return nil
	}
	vals := make([]interface{}, len(cols))
	for i := 0; i < len(cols); i++ {
		vals[i] = new(interface{})
		if i != 0 {
			fmt.Print("\t")
		}
		fmt.Print(cols[i])
	}
	fmt.Println()
	for rows.Next() {
		err = rows.Scan(vals...)
		if err != nil {
			fmt.Println(err)
			continue
		}
		for i := 0; i < len(vals); i++ {
			if i != 0 {
				fmt.Print("\t")
			}
			printValue(vals[i].(*interface{}))
		}
		fmt.Println()

	}
	if rows.Err() != nil {
		return rows.Err()
	}
	return nil
}

func printValue(pval *interface{}) {
	switch v := (*pval).(type) {
	case nil:
		fmt.Print("NULL")
	case bool:
		if v {
			fmt.Print("1")
		} else {
			fmt.Print("0")
		}
	case []byte:
		fmt.Print(string(v))
	case time.Time:
		fmt.Print(v.Format("02-01-2006 15:04:05.999"))
	default:
		fmt.Print(v)
	}
}
