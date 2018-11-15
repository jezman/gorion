package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var disableCmd = &cobra.Command{
	Use:   "disable",
	Short: "disable worker card",
	Run: func(cmd *cobra.Command, args []string) {
		if err := Env.DisableWorkerCard(worker); err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("%s card disabled\n", strings.Title(worker))
		}
	},
}

func init() {
	rootCmd.AddCommand(disableCmd)

	disableCmd.Flags().StringVarP(&worker, "worker", "w", "", "worker first, mid and last name.Use: 'gorion list workers' to get a list of all workers.")
}
