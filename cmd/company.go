package cmd

import (
	"fmt"

	"github.com/jezman/gorion/render"
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

		if companyName != "" {
			employees, err := env.Employees(companyName)
			if err != nil {
				fmt.Println(err)
			}

			table = render.Preparing(employees, "#", "Employee", "Company")
		} else {
			companies, err := env.Company()
			if err != nil {
				fmt.Println(err)
			}

			table = render.Preparing(companies, "#", "Company", "Employees")
		}

		fmt.Println(table.Render())
	},
}

func init() {
	companyCmd.Flags().StringVarP(&companyName, "company", "c", "", "company name")
}
