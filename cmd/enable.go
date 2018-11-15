package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)
var enableCmd = &cobra.Command{
	Use:   "enable",
	Short: "enable worker card",
	Run: func(cmd *cobra.Command, args []string) {
		if err := Env.DisableWorkerCard(worker); err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("%s card enableed\n", worker)
		}
	},
}

func init() {
	rootCmd.AddCommand(enableCmd)

	enableCmd.Flags().StringVarP(&worker, "worker", "w", "", "worker first, mid and last name.Use: 'gorion list workers' to get a list of all workers.")
}
