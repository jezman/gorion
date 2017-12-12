package cmd

import (
	"fmt"

	"github.com/apcera/termtables"
	"github.com/jezman/gorion/query"
	"github.com/spf13/cobra"
)

// eventsCmd represents the events command
var hoursCmd = &cobra.Command{
	Use:   "hours",
	Short: "Displays employees worked time",
	Run: func(cmd *cobra.Command, args []string) {
		db := initDB()
		defer db.Close()

		query := query.WorkedTime(employee, firstDate, lastDate)
		events, err := database.WorkedTime(query)
		if err != nil {
			fmt.Println(err)
		}

		table := termtables.CreateTable()
		table.AddHeaders("Employee", "Company", "First event", "Last event", "Worked time")

		for _, e := range events {
			table.AddRow(
				e.Employee.FullName,
				e.Employee.Company.Name,
				e.FirstTime.Format("02-01-2006 15:04:05"),
				e.LastTime.Format("02-01-2006 15:04:05"),
				e.WorkedTime,
			)
		}
		fmt.Println(table.Render())
	},
}

func init() {
	rootCmd.AddCommand(hoursCmd)

	hoursCmd.Flags().StringVarP(&employee, "employee", "e", "", "employee last name. Use: 'gorion list employees' to get a list of all employees.")
	hoursCmd.Flags().StringVarP(&firstDate, "first", "f", timeNow.Format("02.01.2006"), "first date")
	hoursCmd.Flags().StringVarP(&lastDate, "last", "l", timeNow.AddDate(0, 0, 1).Format("02.01.2006"), "last date.")
}
