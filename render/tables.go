package render

import (
	"github.com/apcera/termtables"
	"github.com/jezman/gorion/helpers"
	"github.com/jezman/gorion/models"
)

// Preparing data for output
func Preparing(data interface{}, headers ...interface{}) *termtables.Table {
	table := termtables.CreateTable()

	for _, header := range headers {
		table.AddHeaders(header)
	}

	switch data.(type) {
	case []*models.Door:
		doors := data.([]*models.Door)
		for _, d := range doors {
			table.AddRow(d.ID, d.Name)
		}
	case []*models.Company:
		companies := data.([]*models.Company)
		for i, c := range companies {
			table.AddRow(i+1, c.Name, c.CountOfEmployees)
		}
	case []*models.Employee:
		employees := data.([]*models.Employee)
		for i, e := range employees {
			table.AddRow(i+1, e.FullName, e.Company.Name)
		}
	case []*models.Event:
		events := data.([]*models.Event)
		for _, e := range events {
			switch {
			case e.WorkedTime > 0:
				table.AddRow(
					e.Employee.FullName,
					e.Employee.Company.Name,
					e.FirstTime.Format("02-01-2006 15:04:05"),
					e.LastTime.Format("02-01-2006 15:04:05"),
					e.WorkedTime,
				)
			case e.Description != "" && e.ID != "":
				table.AddRow(e.ID, e.Action, e.Description)
			case e.Door.Name != "":
				table.AddRow(
					e.FirstTime.Format("15:04:05 02-01-2006"),
					e.Employee.FullName,
					e.Employee.Company.Name,
					e.Door.Name,
					helpers.ColorizedDenied(e.Action),
				)
			}
		}
	}

	return table
}
