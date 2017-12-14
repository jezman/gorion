package cmd

import (
	"fmt"

	"github.com/apcera/termtables"
	"github.com/spf13/cobra"
)

// employeesCmd represents the employees command
var employeesCmd = &cobra.Command{
	Use:     "employees",
	Aliases: []string{"e"},
	Short:   "Displays a list of employees",
	Run: func(cmd *cobra.Command, args []string) {
		db := initDB()
		defer db.Close()

		query := `SELECT p.Name, p.FirstName, p.MidName, c.Name FROM pList p
JOIN pCompany c ON (c.ID = p.Company)
ORDER BY c.Name`

		employees, err := database.Employees(query)
		if err != nil {
			fmt.Println(err)
		}

		table := termtables.CreateTable()
		table.AddHeaders("#", "Company", "Employee")

		for i, e := range employees {
			table.AddRow(i+1, e.FullName, e.Company.Name)
		}

		fmt.Println(table.Render())
	},
}
