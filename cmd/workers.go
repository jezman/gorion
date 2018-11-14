package cmd

import (
	"fmt"

	"github.com/jezman/gorion/render"
	"github.com/spf13/cobra"
)

// workerCmd represents the workers command
var workerCmd = &cobra.Command{
	Use:     "worker",
	Aliases: []string{"w"},
	Short:   "Displays a list of workers",
	Run: func(cmd *cobra.Command, args []string) {
		workers, err := Env.Workers(companyName)
		if err != nil {
			fmt.Println(err)
		}

		table = render.Preparing(workers, "#", "Worker", "Company")
		fmt.Println(table.Render())
	},
}

func init() {
	workerCmd.Flags().StringVarP(&companyName, "company", "c", "", "company name")
}
