package cmd

import (
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l"},
	Short:   "Get list of company, doors, employees",
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.AddCommand(employeesCmd, doorsCmd, companyCmd)
}
