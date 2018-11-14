package cmd

import (
	"fmt"

	"github.com/jezman/gorion/render"
	"github.com/spf13/cobra"
)

var eventsTypes bool

// eventsCmd represents the events command
var eventsCmd = &cobra.Command{
	Use:     "events",
	Aliases: []string{"e"},
	Example: `  gorion events
  gorion events --worker=lastname --first=05.08.2017
  gorion e -e lastname -d 32
  gorion e -d 2 -f 12.11.2017 -l 16.11.2107
  gorion e -f '12.11.2017 21:00'`,
	Short: "Displays a list of events depending on entered flags",
	Run: func(cmd *cobra.Command, args []string) {
		if eventsTypes {
			values, err := Env.EventsValues()

			if err != nil {
				fmt.Println(err)
			}

			table = render.Preparing(values, "ID", "Value", "Description")

		} else {
			events, err := Env.Events(firstDate, lastDate, worker, door, denied)
			if err != nil {
				fmt.Println(err)
			}

			table = render.Preparing(events, "Time", "Worker", "Company", "Door", "Event")
		}

		fmt.Println(table.Render())
	},
}

func init() {
	rootCmd.AddCommand(eventsCmd)
	eventsCmd.AddCommand(tailCmd)

	eventsCmd.Flags().StringVarP(&worker, "worker", "w", "", "worker last name. Use: 'gorion list workers' to get a list of all workers.")
	eventsCmd.Flags().UintVarP(&door, "door", "d", 0, "door ID. Use: 'gorion list doors' to get a list of all doors with ID.")
	eventsCmd.Flags().StringVarP(&firstDate, "first", "f", timeNow.Format("02.01.2006"), "first date")
	eventsCmd.Flags().StringVarP(&lastDate, "last", "l", timeNow.AddDate(0, 0, 1).Format("02.01.2006"), "last date.")
	eventsCmd.Flags().BoolVarP(&eventsTypes, "type", "t", false, "list of events types")
	eventsCmd.Flags().BoolVarP(&denied, "denied", "D", false, "show only denied events")
}
