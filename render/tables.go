package render

import (
	"github.com/apcera/termtables"
	"github.com/jezman/libgorion"
)

// Preparing data for output
func Preparing(data interface{}, headers ...interface{}) *termtables.Table {
	table := termtables.CreateTable()

	for _, header := range headers {
		table.AddHeaders(header)
	}

	switch data.(type) {
	case []*libgorion.Door:
		doors := data.([]*libgorion.Door)
		for _, d := range doors {
			table.AddRow(d.ID, d.Name)
		}
	case []*libgorion.Company:
		companies := data.([]*libgorion.Company)
		for i, c := range companies {
			table.AddRow(i+1, c.Name, c.WorkersCount)
		}
	case []*libgorion.Worker:
		workers := data.([]*libgorion.Worker)
		for i, w := range workers {
			table.AddRow(i+1, w.FullName, w.Company.Name)
		}
	case []*libgorion.Event:
		events := data.([]*libgorion.Event)
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
					libgorion.ColorizedDenied(w.Action),
				)
			}
		}
	}

	return table
}
