package cmd

import (
	"fmt"

	"github.com/apcera/termtables"
	"github.com/spf13/cobra"
)

// companyCmd represents the company command
var companyCmd = &cobra.Command{
	Use:     "company",
	Aliases: []string{"c"},
	Short:   "Displays a list of companies",
	Run: func(cmd *cobra.Command, args []string) {
		db := initDB()
		defer db.Close()

		query := "SELECT Name FROM pCompany"
		company, err := env.Company(query)
		if err != nil {
			fmt.Println(err)
		}

		table := termtables.CreateTable()
		table.AddHeaders("#", "Company")

		for i, c := range company {
			table.AddRow(i+1, c.Name)
		}

		fmt.Println(table.Render())
	},
}
