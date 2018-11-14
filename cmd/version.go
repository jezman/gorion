package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "show application version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("gorion vestion %s\n", VERSION)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
