package cmd

import (
	"fmt"

	"github.com/jezman/gorion/render"
	"github.com/spf13/cobra"
)

// eventsCmd represents the events command
var hoursCmd = &cobra.Command{
	Use:     "hours",
	Aliases: []string{"h"},
	Example: `  gorion hours
  gorion hours --worker=lastname --first=05.08.2017 --last=07.08.2017
  gorion h -e lastname
  gorion h -f 12.11.2017 -l 16.11.2107`,
	Short: "Displays workers worked time",
	Run: func(cmd *cobra.Command, args []string) {
		db := initDB()
		defer db.Close()

		events, err := env.WorkedTime(firstDate, lastDate, worker, companyName)
		if err != nil {
			fmt.Println(err)
		}

		table = render.Preparing(events, "Worker", "Company", "First event", "Last event", "Worked time")
		fmt.Println(table.Render())
	},
}

func init() {
	rootCmd.AddCommand(hoursCmd)

	hoursCmd.Flags().StringVarP(&worker, "worker", "w", "", "worker last name. Use: 'gorion list workers' to get a list of all workers.")
	hoursCmd.Flags().StringVarP(&firstDate, "first", "f", timeNow.Format("02.01.2006"), "first date")
	hoursCmd.Flags().StringVarP(&lastDate, "last", "l", timeNow.AddDate(0, 0, 1).Format("02.01.2006"), "last date.")
	hoursCmd.Flags().StringVarP(&companyName, "company", "c", "", "company name")
}
