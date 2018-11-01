package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var interval int

// tailCmd represents the tail command
var tailCmd = &cobra.Command{
	Use:   "tail",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		intervalDuration := time.Duration(interval)
		tick := time.Tick(intervalDuration * time.Second)

		db := initDB()
		defer db.Close()

		for range tick {
			if err := env.EventsTail(intervalDuration); err != nil {
				fmt.Println(err)
			}
		}
	},
}

func init() {
	tailCmd.Flags().IntVarP(&interval, "interval", "i", 5, "sql queries interval")
}
