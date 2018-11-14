package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add worker to access control system",
	Run: func(cmd *cobra.Command, args []string) {
		if err := Env.AddWorker(worker); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("worker successfully added")
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().StringVarP(&worker, "worker", "w", "", "worker full name(first mid last)")
}
