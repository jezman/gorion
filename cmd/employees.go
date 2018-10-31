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

		employees, err := env.Employees(companyName)
		if err != nil {
			fmt.Println(err)
		}

		table := termtables.CreateTable()
		table.AddHeaders("#", "Employee", "Company")

		for i, e := range employees {
			table.AddRow(i+1, e.FullName, e.Company.Name)
		}

		fmt.Println(table.Render())
	},
}

func init() {
	rootCmd.AddCommand(companyCmd)

	employeesCmd.Flags().StringVarP(&companyName, "company", "c", "", "company name")
}
