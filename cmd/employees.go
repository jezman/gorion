package cmd

import (
	"fmt"

	"github.com/jezman/gorion/render"
	"github.com/spf13/cobra"
)

// employeesCmd represents the workers command
var employeesCmd = &cobra.Command{
	Use:     "workers",
	Aliases: []string{"e"},
	Short:   "Displays a list of workers",
	Run: func(cmd *cobra.Command, args []string) {
		db := initDB()
		defer db.Close()

		workers, err := env.Workers(companyName)
		if err != nil {
			fmt.Println(err)
		}

		table = render.Preparing(workers, "#", "Worker", "Company")
		fmt.Println(table.Render())
	},
}

func init() {
	employeesCmd.Flags().StringVarP(&companyName, "company", "c", "", "company name")
}
