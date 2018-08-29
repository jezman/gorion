package cmd

import (
	"fmt"

	"github.com/apcera/termtables"
	"github.com/spf13/cobra"
)

var companyName string

// companyCmd represents the company command
var companyCmd = &cobra.Command{
	Use:     "company",
	Aliases: []string{"c"},
	Short:   "Displays a list of companies",
	Run: func(cmd *cobra.Command, args []string) {
		db := initDB()
		defer db.Close()

		var table *termtables.Table

		if companyName != "" {
			employees, err := env.Employees(companyName)
			if err != nil {
				fmt.Println(err)
			}

			table = termtables.CreateTable()
			table.AddHeaders("#", "Employee", "Company")

			for i, e := range employees {
				table.AddRow(i+1, e.FullName, e.Company.Name)
			}

		} else {
			company, err := env.Company()
			if err != nil {
				fmt.Println(err)
			}

			table = termtables.CreateTable()
			table.AddHeaders("#", "Company", "Employees")

			for i, c := range company {
				table.AddRow(i+1, c.Name, c.CountOfEmployees)
			}
		}

		fmt.Println(table.Render())
	},
}

func init() {
	rootCmd.AddCommand(companyCmd)

	companyCmd.Flags().StringVarP(&companyName, "company", "c", "", "company name")
}
