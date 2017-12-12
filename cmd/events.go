package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// eventsCmd represents the events command
var eventsCmd = &cobra.Command{
	Use:   "events",
	Short: "Displays a list of events depending on entered flags",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("events called")
	},
}

func init() {
	rootCmd.AddCommand(eventsCmd)

	eventsCmd.Flags().StringVarP(&employee, "employee", "e", "", "employee last name. Use: 'gorion list employees' to get a list of all employees.")
	eventsCmd.Flags().UintVarP(&door, "door", "d", 0, "door ID. Use: 'gorion list doors' to get a list of all doors with ID.")
	eventsCmd.Flags().StringVarP(&firstDate, "first", "f", timeNow.Format("02.01.2006"), "first date")
	eventsCmd.Flags().StringVarP(&lastDate, "last", "l", timeNow.AddDate(0, 0, 1).Format("02.01.2006"), "last date.")
}
