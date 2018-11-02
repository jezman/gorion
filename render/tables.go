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
	case []*models.Worker:
		workers := data.([]*models.Worker)
		for i, w := range workers {
			table.AddRow(i+1, w.FullName, w.Company.Name)
		}
	case []*models.Event:
		events := data.([]*models.Event)
		for _, w := range events {
			switch {
			case w.WorkedTime > 0:
				table.AddRow(
					w.Worker.FullName,
					w.Worker.Company.Name,
					w.FirstTime.Format("02-01-2006 15:04:05"),
					w.LastTime.Format("02-01-2006 15:04:05"),
					w.WorkedTime,
				)
			case w.Description != "" && w.ID != "":
				table.AddRow(w.ID, w.Action, w.Description)
			case w.Door.Name != "":
				table.AddRow(
					w.FirstTime.Format("15:04:05 02-01-2006"),
					w.Worker.FullName,
					w.Worker.Company.Name,
					w.Door.Name,
					helpers.ColorizedDenied(w.Action),
				)
			}
		}
	}

	return table
}
