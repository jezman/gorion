package cmd

import (
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l"},
	Short:   "Get list of company, doors, workers",
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.AddCommand(workerCmd, doorsCmd, companyCmd)
}
