package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/urfave/cli"
)

var (
	db        *sql.DB
	err       error
	rows      *sql.Rows
	user      string
	doorID    string
	firstDate string
	lastDate  string
)

type door struct {
	ID   int64
	Name string
}

type event struct {
	LastName  string
	FirstName string
	MidName   string
	Company   string
	FirstTime time.Time
	LastTime  time.Time
	Events    string
	Door      string
}

type config struct {
	Server   string
	Database string
	User     string
	Password string
}

func readConfigFile() config {
	confFile, err := os.Open("config.json")
	if err != nil {
		fmt.Println("Read configuration file error:", err.Error())
	}
	defer confFile.Close()

	decoder := json.NewDecoder(confFile)
	conf := config{}
	err = decoder.Decode(&conf)
	if err != nil {
		fmt.Println("Json decode error:", err.Error())
	}
	return conf
}

func validationUser(value string) {
	if match, _ := regexp.MatchString(`^[а-яА-Я][а-яА-Я-_\.]{2,20}$`, user); !match {
		os.Exit(1)
	}

}

func validationDate(value string) {
	if match, _ := regexp.MatchString(`(0[1-9]|[12][0-9]|3[01])[- ..](0[1-9]|1[012])[- ..][201]\d\d\d`, value); !match {
		os.Exit(1)
	}
}

func executeQuery(query string) error {
	var conf = readConfigFile()

	dsn := "server=" + conf.Server +
		";user id=" + conf.User +
		";password=" + conf.Password +
		";database=" + conf.Database

	if db, err = sql.Open("mssql", dsn); err != nil {
		fmt.Println("Cannot connect", err.Error())

	}

	if err = db.Ping(); err != nil {
		fmt.Println("Cannot connect", err.Error())
	}
	defer db.Close()

	if rows, err = db.Query(query); err != nil {
		fmt.Println("Query error", err.Error())
	}
	defer rows.Close()

	cols, err := rows.Columns()
	if err != nil {
		fmt.Println("Cols error", err.Error())

	}

	if cols == nil {
		return nil
	}

	vals := make([]interface{}, len(cols))
	for i, col := range cols {
		vals[i] = new(sql.RawBytes)
		fmt.Printf("%-20s", col)
	}
	fmt.Println()

	row := func(cmd ...interface{}) {
		if err := rows.Scan(cmd...); err != nil {
			fmt.Println("Cols error", err.Error())
		}
	}

	for rows.Next() {
		door := door{}
		event := event{}
		switch {
		case len(cols) == 2:
			row(&door.ID, &door.Name)
			fmt.Printf("%-4d %s\n", door.ID, door.Name)
		case len(cols) == 6:
			row(
				&event.LastName,
				&event.FirstName,
				&event.MidName,
				&event.Company,
				&event.FirstTime,
				&event.LastTime,
			)

			// Отработанное время
			diff := event.LastTime.Sub(event.FirstTime)

			fmt.Printf("%-15s %-15s %-15s %-10s %-25s %-25s %s\n",
				event.LastName,
				event.FirstName,
				event.MidName,
				event.Company,
				event.FirstTime.Format("02-01-2006 15:04:05"),
				event.LastTime.Format("02-01-2006 15:04:05"),
				diff,
			)
		case len(cols) == 7:
			row(
				&event.LastName,
				&event.FirstName,
				&event.MidName,
				&event.Company,
				&event.FirstTime,
				&event.Events,
				&event.Door,
			)
			fmt.Printf("%-15s %-15s %-15s %-10s %-25s %-25s %-30s\n",
				event.LastName,
				event.FirstName,
				event.MidName,
				event.Company,
				event.FirstTime.Format("02-01-2006 15:04:05"),
				event.Events,
				event.Door,
			)
		}
	}
	return nil
}

func hours() {
	query := []string{"SELECT p.Name AS Фамилия, p.FirstName AS Имя, p.MidName AS Отчество, c.Name as Компания, min(TimeVal) AS Приход, max(TimeVal) AS Уход ",
		"FROM dbo.pLogData l ",
		"JOIN dbo.pList p ON (p.ID = l.HozOrgan) ",
		"JOIN dbo.pCompany c ON (c.ID = p.Company) ",
		"WHERE TimeVal BETWEEN '", firstDate, "' AND '", lastDate, "'",
		" AND p.Name = '", user, "'",
		" GROUP BY p.Name, p.FirstName, p.MidName, c.Name, CONVERT(varchar(20), TimeVal, 104)",
	}

	validationUser(user)
	validationDate(firstDate)
	validationDate(lastDate)

	if user == "" {
		query = append(query[:9], query[12])
	}

	executeQuery(strings.Join(query, ""))
}

func summary() {
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

	validationUser(user)
	validationDate(firstDate)
	validationDate(lastDate)

	pName := " AND p.Name = '"
	doorIndex := "' AND DoorIndex = "
	orderBy := "' ORDER BY TimeVal"

	add := func(cmd ...string) {
		query = append(query[:len(query)-1], cmd...)
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
	cli.AppHelpTemplate = `ИМЯ:
   {{.Name}} - {{.Usage}}
ИСПОЛЬЗОВАНИЕ:
   {{.HelpName}} {{if .VisibleFlags}}[глобальные параметры]{{end}}{{if .Commands}} команда {{end}}
   {{if len .Authors}}
AUTHOR:
   {{range .Authors}}{{ . }}{{end}}
   {{end}}{{if .Commands}}
КОМАНДЫ:
{{range .Commands}}{{if not .HideHelp}}   {{join .Names ", "}}{{ "\t"}}{{.Usage}}{{ "\n" }}{{end}}{{end}}{{end}}{{if .VisibleFlags}}
ГЛОБАЛЬНЫЕ ПАРАМЕТРЫ:
   {{range .VisibleFlags}}{{.}}
   {{end}}{{end}}
`
	// Значения по умолчанию для первой и последней даты
	timeNow := time.Now().Local()
	firstHourOfDay := timeNow.Format("02.01.2006")
	lastHourOfDay := timeNow.AddDate(0, 0, 1).Format("02.01.2006")

	app := cli.NewApp()
	app.Name = "gorion"
	app.Usage = "создает отчеты для системы контроля доступом НВП Болид \"Орион ПРО\""
	app.HideVersion = true

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "employee, e",
			Value:       "",
			Usage:       "фамилия сотрудника",
			Destination: &user,
		},
		cli.StringFlag{
			Name:        "door, d",
			Value:       "",
			Usage:       "id двери. для просмотра списка всех дверей введите: gorion ld",
			Destination: &doorID,
		},
		cli.StringFlag{
			Name:        "first, f",
			Value:       firstHourOfDay,
			Usage:       "первая дата",
			Destination: &firstDate,
		},
		cli.StringFlag{
			Name:        "last, l",
			Value:       lastHourOfDay,
			Usage:       "последняя дата",
			Destination: &lastDate,
		},
	}

	app.Commands = []cli.Command{
		{
			Name:    "hours",
			Aliases: []string{"h"},
			Usage:   "приход и уход сотрудников + их отработанное время",
			Action: func(c *cli.Context) error {
				hours()
				return nil
			},
		},
		{
			Name:    "summary",
			Aliases: []string{"s"},
			Usage:   "общий отчет",
			Action: func(c *cli.Context) error {
				summary()
				return nil
			},
		},
		{
			Name:    "listdoors",
			Aliases: []string{"ld"},
			Usage:   "список всех дверей с id",
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
