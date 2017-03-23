package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/urfave/cli"
)

var (
	user      string
	doorID    string
	firstDate string
	lastDate  string
)

type door struct {
	ID   int64
	Name string
}

type coWorker struct {
	LastName  string
	FirstName string
	MidName   string
	Company   string
	Time      time.Time
	Events    string
	Door      string
}

type config struct {
	Server   string
	Database string
	User     string
	Password string
}

func checkError(str string, err error) {
	if err != nil {
		fmt.Println(str, err.Error())
	}
}

func readConfigFile() config {
	confFile, err := os.Open("config.json")
	checkError("Read configuration file error:", err)
	defer confFile.Close()

	decoder := json.NewDecoder(confFile)
	conf := config{}
	err = decoder.Decode(&conf)
	checkError("Json decode error:", err)
	return conf
}

func executeQuery(query string) error {
	var conf = readConfigFile()

	dsn := "server=" + conf.Server + ";user id=" + conf.User + ";password=" + conf.Password + ";database=" + conf.Database
	db, err := sql.Open("mssql", dsn)
	checkError("Cannot connect: ", err)

	err = db.Ping()
	checkError("Cannot connect: ", err)
	defer db.Close()
	rows, err := db.Query(query)
	checkError("Query:", err)
	defer rows.Close()

	cols, err := rows.Columns()
	checkError("Cols:", err)

	if cols == nil {
		return nil
	}

	vals := make([]interface{}, len(cols))
	for i, col := range cols {
		vals[i] = new(sql.RawBytes)
		fmt.Printf("%-20s", col)
	}
	fmt.Println()

	if len(cols) == 2 {
		for rows.Next() {
			d := door{}
			err := rows.Scan(&d.ID, &d.Name)
			checkError("Cols:", err)
			fmt.Printf("%-4d %s\n", d.ID, d.Name)
		}
	} else {
		//coWorkers := []coWorker{}
		for rows.Next() {
			mem := coWorker{}
			err := rows.Scan(&mem.LastName, &mem.FirstName, &mem.MidName, &mem.Company, &mem.Time, &mem.Events, &mem.Door)
			checkError("Cols:", err)
			//coWorkers = append(coWorkers, mem)
			fmt.Printf("%-15s %-15s %-15s %-10s %-25s %-25s %-30s\n", mem.LastName, mem.FirstName, mem.MidName, mem.Company, mem.Time.Format("02-01-2006 15:04:05"), mem.Events, mem.Door)
		}
	}
	return nil
}

func generate() {
	query := []string{"SELECT p.Name AS Фамилия, p.FirstName AS Имя, p.MidName AS Отчество, c.Name as Компания, TimeVal AS Время, e.Contents AS Событие, a.Name AS Дверь ",
		"FROM dbo.pLogData l ",
		"JOIN dbo.pList p ON (p.ID = l.HozOrgan) ",
		"JOIN dbo.pCompany c ON (c.ID = p.Company) ",
		"JOIN dbo.Events e ON (e.Event = l.Event) ",
		"JOIN dbo.AcessPoint a ON (a.GIndex = l.DoorIndex) ",
		"WHERE TimeVal BETWEEN '", firstDate, "' AND '", lastDate, "'",
		" AND e.Event BETWEEN 26 AND 29",
		"ORDER BY TimeVal",
	}

	pName := " AND p.Name = '"
	doorIndex := "' AND DoorIndex = "
	orderBy := "' ORDER BY TimeVal"

	add := func(cmd ...string) {
		query = append(query[0:len(query)-1], cmd...)
	}

	if doorID != "" && user != "" {
		add(pName, user, doorIndex, doorID, orderBy[1:])
	} else if user != "" {
		add(pName, user, orderBy)
	} else if doorID != "" {
		add(doorIndex[1:], doorID, orderBy[1:])
	}
	executeQuery(strings.Join(query, ""))

}

func main() {
	// Default first and last date + time
	timeNow := time.Now().Local()
	firstHourOfDay := timeNow.Format("02.01.2006") + " 00:00"
	lastHourOfDay := timeNow.Format("02.01.2006") + " 23:59"

	app := cli.NewApp()

	app.Name = "OrionCLI"
	//app.HelpName = "contrive"
	app.Usage = "generates a reports for Bolid access control system \"Orion Pro\""
	app.UsageText = "orion [global options] command"
	app.HideVersion = true

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "user, u",
			Value:       "",
			Usage:       "user last name",
			Destination: &user,
		},
		cli.StringFlag{
			Name:        "door, d",
			Value:       "",
			Usage:       "door index. For show all doors indexes use: orion ld",
			Destination: &doorID,
		},
		cli.StringFlag{
			Name:        "first, f",
			Value:       firstHourOfDay,
			Usage:       "first date of a report",
			Destination: &firstDate,
		},
		cli.StringFlag{
			Name:        "last, l",
			Value:       lastHourOfDay,
			Usage:       "last date of a report",
			Destination: &lastDate,
		},
	}

	app.Commands = []cli.Command{
		{
			Name:    "generate",
			Aliases: []string{"g"},
			Usage:   "generate a report",
			Action: func(c *cli.Context) error {
				generate()
				return nil
			},
		},
		{
			Name:    "listdoors",
			Aliases: []string{"ld"},
			Usage:   "list all doors with indexes",
			Action: func(c *cli.Context) error {
				executeQuery("SELECT GIndex as ID, Name as Дверь from dbo.AcessPoint")
				return nil
			},
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	app.Run(os.Args)

}
